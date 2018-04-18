/* file: $GOPATH/src/buildstack/credhub_test.go */
package main

import (
    "bytes"
    "fmt"
    "github.com/DATA-DOG/godog"
    "net/url"
    "os"
    "os/exec"
    "strings"
)

var rawCredHubUrl string
var secretPath string
var usernameVariable string
var passwordVariable string

//http://craigwickesser.com/2015/02/golang-cmd-with-custom-environment/
//func ExecViaBoshProxy(command *string, args *[]string) (string, error) {
//    var (
//        commandOut []byte
//        commandStdout bytes.Buffer
//        error error
//    )
//
//    commandExec = exec.Command(*command, *args...).CombinedOutput()
//    environ := os.Environ()
//    environ = append(environ, "https_proxy=${BOSH_ALL_PROXY}")
//    commandExec.Env = environ
//    commandExec.Stdin = bytes.NewBuffer(commandOut)
//    commandExec.Stdout = &commandStdout
//    error = command.Run()
//}

func LoginToCredHub() error {
    var (
        commandOut []byte
        error error
    )

    credHubCommand := "credhub"
    apiArgs := []string{"api", "--server=" + rawCredHubUrl}
//    fmt.Fprint(os.Stdout, "Setting CredHub server to be ", rawCredHubUrl, "\n")
    commandOut, error = exec.Command(credHubCommand, apiArgs...).CombinedOutput()

    if error != nil {
        return fmt.Errorf("There was an error running the credhub api command: ", error, "\nOutput: ", string(commandOut[:]))
    }

    loginArgs := []string{"login", "--username=" + os.Getenv(usernameVariable), "--password=" + os.Getenv(passwordVariable)}
//    fmt.Fprint(os.Stdout, "Logging in using username (credHubUsernameVariable) ", os.Getenv(usernameVariable), " and password (credHubPasswordVariable) ", os.Getenv(passwordVariable), "\n")
    commandOut, error = exec.Command(credHubCommand, loginArgs...).CombinedOutput()

    if error != nil {
        return fmt.Errorf("There was an error running the credhub login command: ", error, "\nOutput: ", string(commandOut[:]))
    }

    return nil
}

func LogoutOfCredHub() error {
    var (
        commandOut []byte
        error error
    )

    credHubCommand := "credhub"
    logoutArgs := []string{"logout"}
    commandOut, error = exec.Command(credHubCommand, logoutArgs...).CombinedOutput()

    if error != nil {
        return fmt.Errorf("There was an error running the credhub logout command: ", error, "\nOutput: ", string(commandOut[:]))
    }

    return nil
}

// Pointers didn't work.
func CredHubIsInstalled(credHubUrl, credHubUsernameVariable, credHubPasswordVariable string) error {
    rawCredHubUrl = credHubUrl
    usernameVariable = credHubUsernameVariable
    passwordVariable = credHubPasswordVariable

    _, error := url.ParseRequestURI(rawCredHubUrl)

    if error != nil {
        return fmt.Errorf("CredHub URL %v is invalid", rawCredHubUrl)
    }

    return nil
}

func ASecretExists(expectedSecretValue, credHubSecretPath string) error {
    if credHubSecretPath == "" { // Should be enforced by the regex anyway
        return fmt.Errorf("A secret path must be specified")
    }

    secretPath = credHubSecretPath

    error := LoginToCredHub()

    if error != nil {
        return error
    }

    var (
        commandOut []byte
    )
//credHubApiBaseContextPath 
    credHubCommand := "credhub"
    setArgs := []string{"set", "--name=" + secretPath, "--value=" + expectedSecretValue, "--type=value"}
//    fmt.Fprint(os.Stdout, "Setting secret ", expectedSecretValue, " at path ", secretPath, " using command arguments: ", setArgs, "\n")
    commandOut, error = exec.Command(credHubCommand, setArgs...).CombinedOutput()

    LogoutOfCredHub()

    if error != nil {
        return fmt.Errorf("There was an error running the credhub set command: ", error, "\nOutput: ", string(commandOut[:]))
    }

    return nil
}

// Could be done within the container setup instead?
func AccessCredHub() error {
    LoginToCredHub()

    return nil
}

func SecretValueShouldBe(expectedSecretValue string) error {
    var (
        credHubCommandOut []byte
        jqCommandOut []byte
        jqStdout bytes.Buffer
        error error
    )

    credHubCommand := "credhub"
    getArgs := []string{"get", "--name=" + secretPath,  "--output-json"}
//    fmt.Fprintln(os.Stdout, "Getting secret from path ", secretPath)
    credHubCommandOut, error = exec.Command(credHubCommand, getArgs...).CombinedOutput()

    if error != nil {
        return fmt.Errorf("There was an error running the credhub get command: ", error, "\nOutput: ", string(credHubCommandOut[:]))
    }

    jqCommand := "jq"
    jqArgs := []string{"-r", ".value"}
    jqCommandExec := exec.Command(jqCommand, jqArgs...)
    jqCommandExec.Stdin = bytes.NewBuffer(credHubCommandOut)
	jqCommandExec.Stdout = &jqStdout
    error = jqCommandExec.Run()

    if error != nil {
        return fmt.Errorf("There was an error running the credhub get jq command: ", error, "\nOutput: ", string(jqCommandOut[:]))
    }

    secretValue := strings.TrimSpace(string(jqStdout.Bytes()))

    fmt.Fprintln(os.Stdout, "The retrieved secret value at " + secretPath + " was " + secretValue)

    if (secretValue != expectedSecretValue) {
        return fmt.Errorf("The expected secret value was " + expectedSecretValue + ", not " + secretValue)
    }

    return nil
}

func CredHubFeatureContext(suite *godog.Suite) {
    suite.Step(`^CredHub is installed at (https:\/\/\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b:[0-9]*) with username \${([_-a-zA-Z0-9]*)} and password \${([_-a-zA-Z0-9]*)}$`, CredHubIsInstalled)
    suite.Step(`^a secret (.*) exists at ([-a-z0-9\/]*)$`, ASecretExists)
    suite.Step(`^I use the CredHub CLI to access CredHub via a proxy$`, AccessCredHub)
    suite.Step(`^the secret value should be (.*)$`, SecretValueShouldBe)

    suite.BeforeScenario(func(interface{}) {
        rawCredHubUrl = ""
        secretPath = ""
        usernameVariable = ""
        passwordVariable = ""
    })

    suite.AfterScenario(func(interface{}, error) {
        LogoutOfCredHub()
    })
}
