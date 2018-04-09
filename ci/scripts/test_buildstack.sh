#!/bin/bash -ex

export TERM=dumb &&\
cp -r godog-test/gosrc/* ${GOPATH} &&\
cd godog-test &&\
./build_and_run.sh ${JUMPBOX_ADDRESS} ${JUMPBOX_PRIVATE_KEY} ${CREDHUB_USERNAME} ${CREDHUB_PASSWORD} ${CREDHUB_PROXY_PORT} ${https_proxy}
