#!/bin/bash -e

export TERM=dumb &&\
cp -r godog-test/gosrc/* ${GOPATH} &&\
cd godog-test &&\
echo "${JUMPBOX_PRIVATE_KEY}" > /tmp/jumpbox.key
./build_and_run.sh "${JUMPBOX_ADDRESS}" /tmp/jumpbox.key "${CREDHUB_USERNAME}" "${CREDHUB_PASSWORD}" "${CREDHUB_PROXY_PORT}"
rm > /tmp/jumpbox.key
