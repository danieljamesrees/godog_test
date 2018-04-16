#!/bin/sh -ex

DEFAULT_CREDHUB_PROXY_PORT="6666"

JUMPBOX_ADDRESS="${1}"
JUMPBOX_PRIVATE_KEY="${2}"
CREDHUB_USERNAME="${3}"
CREDHUB_PASSWORD="${4}"
CREDHUB_PROXY_PORT="${5}"

usage()
{
    echo $0 JUMPBOX_ADDRESS JUMPBOX_PRIVATE_KEY CREDHUB_USERNAME CREDHUB_PASSWORD [CREDHUB_PROXY_PORT]
}

if [ -z "${JUMPBOX_ADDRESS}" ]
then
    echo Must specify a jumpbox address
    usage
    exit 1
fi

if [ -z "${JUMPBOX_PRIVATE_KEY}" ]
then
    echo Must specify a jumpbox private key
    usage
    exit 1
fi

if [ -z "${CREDHUB_USERNAME}" ]
then
    echo Must specify a CredHub username
    usage
    exit 1
fi

if [ -z "${CREDHUB_PASSWORD}" ]
then
    echo Must specify a CredHub password
    usage
    exit 1
fi

if [ -z "${CREDHUB_PROXY_PORT}" ]
then
    CREDHUB_PROXY_PORT="${DEFAULT_CREDHUB_PROXY_PORT}"
fi

clean()
{
    rm -rf gosrc/bin/tests
#    rm --force --preserve-root --recursive gosrc/bin/tests
}

init()
{
    if [ "${GOPATH}" != "/go" ]
    then
        export GOPATH="${PWD}/gosrc"
    fi

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

    if [ ! -f /usr/local/bin/credhub ] # --version requires a server connection, when already configured
    then
        echo Installing credhub CLI
        curl --location --output credhub.tgz https://github.com/cloudfoundry-incubator/credhub-cli/releases/download/1.7.0/credhub-linux-1.7.0.tgz
        tar -xvf credhub.tgz
        rm -f credhub.tgz
#        rm --force credhub.tgz
        chmod u+x credhub
        sudo mv credhub /usr/local/bin/
    fi

    mkdir -p bin/tests
#    mkdir --parents bin/tests
}

build()
{
    godog --format=cucumber --output bin/tests/credhub --strict ${PWD}/src/buildstack/features/credhub.feature
    godog --format=cucumber --output bin/tests/jenkins --strict ${PWD}/src/buildstack/features/jenkins.feature
}

# Hopefully this proxy configuration can eventually be removed if CredHub can be made accessible from Concourse.
setup_credhub()
{
    # Will fail if the port is already open.
#    ssh -o StrictHostKeyChecking=no -N -D ${CREDHUB_PROXY_PORT} jumpbox@${JUMPBOX_ADDRESS} -i "${JUMPBOX_PRIVATE_KEY}" &
#    export CREDHUB_PROXY=socks5://localhost:${CREDHUB_PROXY_PORT}
    export https_proxy=${BOSH_ALL_PROXY}
}

clean
init
build

setup_credhub

bin/tests/credhub ${PWD}/src/buildstack/features/credhub.feature
bin/tests/jenkins ${PWD}/src/buildstack/features/jenkins.feature
