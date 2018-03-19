#!/bin/bash -e

set -o pipefail

PIPELINE_NAME="${PIPELINE_NAME}"

usage()
{
    echo $0
}

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

TARGET_NAME="concourse-dev"

echo Ensure you are logged into the correct Concourse instance using fly --target ${TARGET_NAME} login --team-name finkit-cpo --concourse-url CONCOURSE_EXTERNAL_URL
echo About to start pipeline - do not quit until the smoking_pipeline Configured message is displayed or errors are identified

fly --target ${TARGET_NAME} set-pipeline --pipeline ${PIPELINE_NAME} --config ci/${PIPELINE_NAME}.yml --non-interactive
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

fly --target ${TARGET_NAME} trigger-job --job ${PIPELINE_NAME}/test_jenkins_job --watch
fail_on_error "Failed to trigger test_jenkins job"

echo ${PIPELINE_NAME} Configured
