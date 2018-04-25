#!/bin/bash -e

export TERM=dumb &&\
curl --location --output bbl https://github.com/cloudfoundry/bosh-bootloader/releases/download/v5.11.6/bbl-v5.11.6_linux_x86-64 &&\
chmod u+x bbl &&\
sudo mv bbl /usr/local/bin/ &&\
apk add jq &&\
cd buildstack-bbl-state/buildstack-bbl-state &&\
eval "$(bbl print-env)" &&\
cd ../.. &&\
mkdir -p ~/.ssh &&\
chmod u=rwx,go= ~/.ssh &&\
echo "${JUMPBOX_PRIVATE_KEY}" > ~/.ssh/jumpbox.key &&\
chmod u=rw,go= ~/.ssh/jumpbox.key &&\
cp -r godog-test/gosrc/* ${GOPATH} &&\
cd godog-test &&\
./build_and_run.sh "${JUMPBOX_ADDRESS}" "${CREDHUB_PROXY_PORT}" &&\
rm ~/.ssh/jumpbox.key
