# file: $GOPATH/src/godogs/features/credhub.feature
Feature: access credhub secrets
  In order to securely test BOSH deployments and releases
  As a CPO Engineer
  I need to be able to get credentials from CredHub

  @initialisation
  Scenario: Access CredHub secrets
    Given CredHub is installed at https://10.0.0.6:8844
    And a secret exists at /concourse/main/godog-test/ive-got-the-key
    When I use the CredHub CLI to access CredHub via a proxy with username ${CREDHUB_USERNAME} and password ${CREDHUB_PASSWORD}
    Then the secret value should be ive-got-the-secret
