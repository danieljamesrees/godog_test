# file: $GOPATH/src/godogs/features/credhub.feature
Feature: access credhub secrets
  In order to securely test BOSH deployments and releases
  As a CPO Engineer
  I need to be able to get credentials from CredHub

  @initialisation
  Scenario: Access CredHub secrets
    Given CredHub is installed at ${CREDHUB_SERVER} with username ${CREDHUB_USER} and password ${CREDHUB_PASSWORD}
    And a secret ive-got-the-secret exists at /godog-test/ive-got-the-key
    When I use the CredHub CLI to access CredHub via a proxy
    Then the secret value should be ive-got-the-secret
