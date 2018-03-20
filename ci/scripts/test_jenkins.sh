#!/bin/bash -ex

export TERM=dumb &&\
mkdir ${GOPATH}/src &&\
cp --archive godog_test/src ${GOPATH}/src &&\
cd ${GOPATH}/godog_test &&\
# go gets?
godog
