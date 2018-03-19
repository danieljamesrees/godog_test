/* file: $GOPATH/src/buildstack/jenkins_test.go */
package main

import (
    "fmt"

    "github.com/DATA-DOG/godog"
)

func jenkinsInstalled(jenkinsUrl string) error {
    pingOutput = ""
    return nil
}

func iPingJenkinsUrl(jenkinsUrl string) error {
    if jenkinsUrl is invalid {
        return fmt.Errorf("Jenkins URL %v is invalid", jenkinsUrl)
    }
    pingOutput = "something"
    return nil
}

func itShouldBeUnlocked error {
    if pingOutput != remaining {
        return fmt.Errorf("Jenkins is still locked according to %v", pingOutput)
    }
    return nil
}

func FeatureContext(suite *godog.Suite) {
    suite.Step(`^Jenkins is installed at (\v+)$`, jenkinsInstalled)
    suite.Step(`^I ping its URL$`, iPingJenkinsURL)
    suite.Step(`^it should be unlocked$`, itShouldBeUnlocked)

    suite.BeforeScenario(func(interface{}) {
        pingOutput = "" // clean the state before every scenario
    })
}
