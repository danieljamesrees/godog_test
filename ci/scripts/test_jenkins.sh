#!/bin/bash -ex

export TERM=dumb &&\
cp -a godog_test/gosrc ${GOPATH} &&\
cd godog_test &&\
./build_and_run.sh
