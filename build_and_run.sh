#!/bin/sh -ex

clean()
{
    rm --force --preserve-root --recursive bin/tests
    rm --force --preserve-root --recursive src/github.com
}

init()
{
    export GOPATH=$(pwd)
    #git clone https://github.com/DATA-DOG/godog.git $GOPATH/src/buildstack/vendor/github.com/DATA-DOG/godog
    go get github.com/DATA-DOG/godog/cmd/godog

    mkdir bin/tests
}

build()
{
    bin/godog --format=cucumber --output bin/tests/jenkins --strict $(pwd)/src/buildstack/features/jenkins.feature
}

clean
init
build

bin/tests/jenkins $(pwd)/src/buildstack/features/jenkins.feature
