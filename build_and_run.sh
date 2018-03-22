#!/bin/sh -ex

# Only needed locally.
clean()
{
    rm -rf gosrc/bin/tests
#    rm --force --preserve-root --recursive gosrc/bin/tests
}

init()
{
    export GOPATH="${PWD}/gosrc"
    cd ${GOPATH}
    export PATH="${GOPATH}/bin:${PATH}"

    set +x
    if ! godog --version
    then
        echo Installing godog
        #git clone https://github.com/DATA-DOG/godog.git $GOPATH/src/buildstack/vendor/github.com/DATA-DOG/godog
        go get github.com/DATA-DOG/godog/cmd/godog
    fi
    set -x

    mkdir -p bin/tests
#    mkdir --parents bin/tests
}

build()
{
    godog --format=cucumber --output bin/tests/jenkins --strict ${PWD}/src/buildstack/features/jenkins.feature
}

clean
init
build

bin/tests/jenkins ${PWD}/src/buildstack/features/jenkins.feature
