/* file: $GOPATH/src/buildstack/jenkins_test.go */
package main

import (
//    "bytes"
    "crypto/tls"
    "fmt"
    "github.com/DATA-DOG/godog"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

var rawJenkinsUrl string
var transport = &http.Transport{}

func JenkinsIsInstalled(jenkinsUrl string) error {
    rawJenkinsUrl = jenkinsUrl
    return nil
}

func CertificatesAreBeingIgnored() error {
    transport = &http.Transport {
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    return nil
}

func IPingJenkins() error {
    _, error := url.ParseRequestURI(rawJenkinsUrl)

    if error != nil {
        return fmt.Errorf("Jenkins URL %v is invalid", rawJenkinsUrl)
    }

    return nil
}

func ItShouldBeUnlocked() error {
    client := &http.Client{Transport: transport}
    jenkinsResponse, error := client.Get(rawJenkinsUrl)

    if error != nil {
        return fmt.Errorf("Failed to contact Jenkins URL %v", rawJenkinsUrl)
    }

    jenkinsResponseBodyBytes, error := ioutil.ReadAll(jenkinsResponse.Body)
    jenkinsResponse.Body.Close()
//    jenkinsResponseBodyLength := bytes.IndexByte(jenkinsResponseBody, 0)
    jenkinsResponseBody := string(jenkinsResponseBodyBytes[:])

    if strings.Contains(jenkinsResponseBody, "Unlock Jenkins") {
//    if jenkinsResponseBodyLength < 1 && !strings.Contains(string(jenkinsResponseBody[:]), "Unlock Jenkins") {
        return fmt.Errorf("Jenkins is still locked according to %v", jenkinsResponseBody)
    }

    return nil
}

func FeatureContext(suite *godog.Suite) {
    suite.Step(`^Jenkins is installed at (https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*))$`, JenkinsIsInstalled)
    suite.Step(`^certificates are being ignored$`, CertificatesAreBeingIgnored)
    suite.Step(`^I ping its URL$`, IPingJenkins)
    suite.Step(`^it should be unlocked$`, ItShouldBeUnlocked)

    suite.BeforeScenario(func(interface{}) {
        rawJenkinsUrl = ""
        transport = nil
    })
}
