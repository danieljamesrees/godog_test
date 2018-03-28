/* file: $GOPATH/src/buildstack/credhub_test.go */
package main

// Ensure HTTPS_PROXY is unset.

import (
    "fmt"
    "github.com/DATA-DOG/godog"
    "net/http"
    "net/url"
    "os"
    "os/exec"
    "strings"
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
    if (secretPath == "") // Should be enforced by the regex anyway
    {
        return fmt.Errorf("A secret path under /concourse must be specified")
    }

    secretPath = secretPath
    return nil
}

// Could be done within the container setup instead?
func AccessCredHub(credHubUsername string, credHubPassword string) error {
    var (
        commandOut []byte
        error    error
    )

    credHubCommand := "credhub"
    apiArgs := []string{"api", "--server", "10.0.0.6"}
    commandOut, error = exec.Command(credHubCommand, apiArgs...).Output()

    if error != nil {
        fmt.Errorf(os.Stderr, "There was an error running the credhub api command: ", error, " - output was ", commandOut)
    }
    else
    {
        loginArgs := []string{"login", "--username", credHubUsername, "--password", credHubPassword} // --server?
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
        fmt.Errorf(os.Stderr, "There was an error running the credhub get command: ", error, " - output was ", commandOut)
    }
	else
	{
        secretValue := string(commandOut)

        fmt.Println("The retrieved secret value at ", secretPath, " was ", secretValue)

        if (secretValue != expectedSecretValue)
        {
            fmt.Errorf("The expected secret value was ", expectedSecretValue, ", not ", secretValue)
        }
    }
}

func FeatureContext(suite *godog.Suite) {
    suite.Step(`^CredHub is installed at (https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*:[0-9]*))$`, CredHubIsInstalled)
    suite.Step(`^a secret exists at /concourse/([-a-z0-9\/]*)$`, ASecretExists)
    suite.Step(`^I use the CredHub CLI to access CredHub via a proxy with the username ([-a-z0-9\/]*) and password ([-a-z0-9\/]*)$`, AccessCredHub)
    suite.Step(`^the secret value should be (.*)$`, SecretValueShouldBe)

    suite.BeforeScenario(func(interface{}) {
        credHubUrl = ""
    })
}
