#!/bin/bash -e

# Typically runs as $0 test-buildstack

set -o pipefail

PIPELINE_NAME="godog-test" # Must match the path in Vault #"$(basename $PWD)"

usage()
{
    echo $0 BBL_STATE_PATH
}

if [ -z "${BBL_STATE_PATH}" ]
then
    echo Must specify a BBL state path
    usage
    exit 1
fi

fail_on_error()
{
    local _message="${1}"
    local _error_code=$?

    if [ $? -ne 0 ]
    then
        echo ${_message}
        usage
        exit ${_error_code}
    fi
}

# TODO Get ssh-key and credhub-password from Vault
setup_credhub()
{
    echo "---" > /tmp/${PIPELINE_NAME}-vars.yml

    echo "jumpbox-address: $(bbl --state-dir ${BBL_STATE_PATH} jumpbox-address)" >> /tmp/${PIPELINE_NAME}-vars.yml
    echo "credhub-proxy-port: 6666" >> /tmp/${PIPELINE_NAME}-vars.yml
}

TARGET_NAME="concourse-dev"

setup_credhub

echo Ensure you are logged into the correct Concourse instance using fly --target ${TARGET_NAME} login --team-name YOUR_TEAM_NAME --concourse-url CONCOURSE_EXTERNAL_URL
echo About to start pipeline - do not quit until the smoking_pipeline Configured message is displayed or errors are identified

fly --target ${TARGET_NAME} set-pipeline --pipeline ${PIPELINE_NAME} --config ci/${PIPELINE_NAME}.yml --non-interactive --load-vars-from=/tmp/${PIPELINE_NAME}-vars.yml
fail_on_error "Failed to set ${PIPELINE_NAME} pipeline"

echo
if ! fly --target ${TARGET_NAME} pipelines|grep --silent ${PIPELINE_NAME}
then
   echo Pipeline not set up - may need to be run a few times after the initial Concourse setup/pipeline destruction
   exit 1
fi
echo Finished setting up pipeline

fly --target ${TARGET_NAME} unpause-pipeline --pipeline ${PIPELINE_NAME}
fail_on_error "Failed to unpause ${PIPELINE_NAME} pipeline"

fly --target ${TARGET_NAME} trigger-job --job ${PIPELINE_NAME}/test-buildstack-job --watch
fail_on_error "Failed to trigger test_jenkins job"

echo ${PIPELINE_NAME} Configured
