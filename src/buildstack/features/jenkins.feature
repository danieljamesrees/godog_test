# file: $GOPATH/src/godogs/features/jenkins.feature
Feature: ping jenkins
  In order to be release Jenkins for use
  As a CPO Engineer
  I need to be able to check Jenkins is ready for use

  Scenario: Jenkins unlocked
    Given Jenkins is installed
    When I ping its URL
    Then it should be unlocked
