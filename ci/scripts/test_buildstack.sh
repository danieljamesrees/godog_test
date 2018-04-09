#!/bin/bash -e

export TERM=dumb &&\
mkdir ~/.ssh &&\
chmod u=rwx,go= ~/.ssh &&\
echo "${JUMPBOX_PRIVATE_KEY}" > ~/.ssh/jumpbox.key &&\
chmod u=rw,go= ~/.ssh/jumpbox.key &&\
cp -r godog-test/gosrc/* ${GOPATH} &&\
cd godog-test &&\
./build_and_run.sh "${JUMPBOX_ADDRESS}" ~/.ssh/jumpbox.key "${CREDHUB_USERNAME}" "${CREDHUB_PASSWORD}" "${CREDHUB_PROXY_PORT}"
rm ~/.ssh/jumpbox.key
