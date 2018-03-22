#!/bin/bash -ex

export TERM=dumb &&\
export GOPATH=godog_test &&\
cd ${GOPATH}/gosrc &&\
./build_and_run.sh
