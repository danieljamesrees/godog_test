#!/bin/bash -ex

export TERM=dumb &&\
cp -r godog_test/gosrc/* ${GOPATH} &&\
cd godog_test &&\
./build_and_run.sh ${JUMPBOX_ADDRESS} ${JUMPBOX_PRIVATE_KEY} ${CREDHUB_USERNAME} ${CREDHUB_PASSWORD} ${CREDHUB_PROXY_PORT} ${https_proxy}
