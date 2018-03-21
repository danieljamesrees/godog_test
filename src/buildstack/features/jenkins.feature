# file: $GOPATH/src/godogs/features/jenkins.feature
Feature: ping jenkins
  In order to be release Jenkins for use
  As a CPO Engineer
  I need to be able to check Jenkins is ready for use

  @initialisation
  Scenario: Jenkins unlocked
    Given Jenkins is installed at https://jenkins.dev-build-create.build.finkit.io
    And certificates are being ignored
    When I ping its URL
    Then it should be unlocked
