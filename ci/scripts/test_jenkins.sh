#!/bin/bash -ex

export TERM=dumb &&\
mkdir ${GOPATH}/src &&\
cp --archive godog_test/src ${GOPATH}/src &&\
cd ${GOPATH}/godog_test &&\
go get github.com/DATA-DOG/godog/cmd/godog
mkdir bin/tests
godog --format=cucumber --output bin/tests/jenkins --strict $(pwd)/src/buildstack/features/jenkins.feature
