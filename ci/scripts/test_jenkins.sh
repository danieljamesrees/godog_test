#!/bin/bash -ex

export TERM=dumb &&\
cp -r godog_test/gosrc/* ${GOPATH} &&\
cd godog_test &&\
./build_and_run.sh
