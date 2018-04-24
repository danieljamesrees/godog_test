#!/bin/bash -e

export TERM=dumb &&\
mkdir ~/.ssh &&\
chmod u=rwx,go= ~/.ssh &&\
echo "${JUMPBOX_PRIVATE_KEY}" > ~/.ssh/jumpbox.key &&\
chmod u=rw,go= ~/.ssh/jumpbox.key &&\
cd buildstack-bbl-state/buildstack-bbl-state &&\
eval "$(bbl print-env)" &&\
cd .. &&\
cp -r godog-test/gosrc/* ${GOPATH} &&\
cd godog-test &&\
./build_and_run.sh "${JUMPBOX_ADDRESS}" "${CREDHUB_PROXY_PORT}"
rm ~/.ssh/jumpbox.key
