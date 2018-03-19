# file: $GOPATH/src/godogs/features/jenkins.feature
Feature: ping jenkins
  In order to be check Jenkins is unlocked
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining
