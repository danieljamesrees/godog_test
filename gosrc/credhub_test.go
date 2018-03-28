/* file: $GOPATH/src/buildstack/credhub_test.go */
package main

// Ensure HTTPS_PROXY is unset.

import (
    "fmt"
    "github.com/DATA-DOG/godog"
    "net/url"
    "os"
    "os/exec"
)

var rawCredHubUrl string
var secretPath string

func CredHubIsInstalled(credHubUrl string) error {
    rawCredHubUrl = credHubUrl

    _, error := url.ParseRequestURI(rawCredHubUrl)

    if error != nil {
        return fmt.Errorf("CredHub URL %v is invalid", rawCredHubUrl)
    }

    return nil
}

func ASecretExists(secretPath string) error {
    if secretPath == "" { // Should be enforced by the regex anyway
        return fmt.Errorf("A secret path under /concourse must be specified")
    }

    secretPath = secretPath
    return nil
}

// Could be done within the container setup instead?
func AccessCredHub(credHubUsernameVariable string, credHubPasswordVariable string) error {
    var (
        commandOut []byte
        error    error
    )

    credHubCommand := "credhub"
    apiArgs := []string{"api", "--server", rawCredHubUrl}
    commandOut, error = exec.Command(credHubCommand, apiArgs...).Output()

    if error != nil {
        fmt.Errorf("There was an error running the credhub api command: ", error, " - output was ", commandOut)
    } else {
        loginArgs := []string{"login", "--username", os.Getenv("credHubUsernameVariable"), "--password", os.Getenv("credHubPasswordVariable")} // --server?
        commandOut, error = exec.Command(credHubCommand, loginArgs...).Output()
    }

    return nil
}

func SecretValueShouldBe(expectedSecretValue string) error {
    var (
        commandOut []byte
        error    error
    )

    credHubCommand := "credhub"
    getArgs := []string{"get", "--name=", secretPath}
    commandOut, error = exec.Command(credHubCommand, getArgs...).Output()

    if error != nil {
        fmt.Errorf("There was an error running the credhub get command: ", error, " - output was ", commandOut)
    } else {
        secretValue := string(commandOut)

        fmt.Println("The retrieved secret value at ", secretPath, " was ", secretValue)

        if (secretValue != expectedSecretValue) {
            fmt.Errorf("The expected secret value was ", expectedSecretValue, ", not ", secretValue)
        }
    }

    return nil
}

func CredHubFeatureContext(suite *godog.Suite) {
    suite.Step(`^CredHub is installed at (https:\/\/\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b:[0-9]*)$`, CredHubIsInstalled)
    suite.Step(`^a secret exists at /concourse/([-a-z0-9\/]*)$`, ASecretExists)
    suite.Step(`^I use the CredHub CLI to access CredHub via a proxy with username \${([_-a-zA-Z0-9]*)} and password \${([_-a-zA-Z0-9]*)}$`, AccessCredHub)
    suite.Step(`^the secret value should be (.*)$`, SecretValueShouldBe)

    suite.BeforeScenario(func(interface{}) {
        rawCredHubUrl = ""
        secretPath = ""
    })
}
