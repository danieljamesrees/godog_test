/* file: $GOPATH/src/buildstack/jenkins_test.go */
package main

import (
    "fmt"
    "github.com/DATA-DOG/godog"
    "io/ioutil"
    "net.url"
    "strings"
)

func JenkinsInstalled(jenkinsUrl string) error {
    rawJenkinsUrl = jenkinsUrl
    return nil
}

func IPingJenkins() error {
     parsedJenkinsUrl, error := url.ParseRequestURI(rawJenkinsUrl)

    if error != nil {
        return fmt.Errorf("Jenkins URL %v is invalid", rawJenkinsUrl)
    }

    return nil
}

func ItShouldBeUnlocked() error {
    jenkinsResponse, error := http.Get(rawJenkinsUrl)

    if error !=nil {
        return fmt.Errorf("Failed to contact Jenkins URL %v", rawJenkinsUrl)
    }

    jenkinsResponseBody, error := ioutil.ReadAll(jenkinsResponse.Body)
    jenkinsResponse.Body.Close()

    if !strings.Contains(jenkinsResponseBody, "Unlock Jenkins") {
        return fmt.Errorf("Jenkins is still locked according to %v", jenkinsResponseBody)
    }

    return nil
}

func FeatureContext(suite *godog.Suite) {
    suite.Step(`^Jenkins is installed at (\v+)$`, JenkinsInstalled)
    suite.Step(`^I ping its URL$`, IPingJenkins)
    suite.Step(`^it should be unlocked$`, ItShouldBeUnlocked)

    suite.BeforeScenario(func(interface{}) {
        jenkinsResponse = "" // clean the state before every scenario
    })
}
