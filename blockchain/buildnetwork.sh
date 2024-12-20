#!/bin/bash

ROOTDIR=$(cd "$(dirname "$0")" && pwd)
export PATH=${ROOTDIR}/bin:${PWD}/bin:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx
export VERBOSE=false

pushd ${ROOTDIR} > /dev/null
trap "popd > /dev/null" EXIT

CONTAINER_CLI=docker
CONTAINER_CLI_COMPOSE=docker-compose
COMPOSE_FILE_CAR=docker-compose-car.yaml
COMPOSE_FILE_CA=docker-compose-ca.yaml
# Get docker sock path from environment variable
SOCK="${DOCKER_HOST:-/var/run/docker.sock}"
DOCKER_SOCK="${SOCK##unix://}"

. scripts/utils.sh
. scripts/createdir.sh
function prereqsInfo() {
    # check docker image
    LOCAL_VERSION=$(peer version | sed -ne 's/^ Version: //p')
    DOCKER_IMAGE_VERSION=$(docker run --rm hyperledger/fabric-peer:latest peer version |
                                 sed -ne 's/^ Version: //p')
    infoln "LOCAL_VERSION=$LOCAL_VERSION"
    infoln "DOCKER_IMAGE_VERSION=$DOCKER_IMAGE_VERSION"
    if [ "$LOCAL_VERSION" != "$DOCKER_IMAGE_VERSION" ]; then
        warnln "Local fabric binaries and docker images are out of  sync. This may cause problems."
    fi

    # check fabric ca
    CA_LOCAL_VERSION=$(fabric-ca-client version | sed -ne 's/ Version: //p')
    CA_DOCKER_IMAGE_VERSION=$(${CONTAINER_CLI} run --rm hyperledger/fabric-ca:latest fabric-ca-client version | sed -ne 's/ Version: //p' | head -1)
    infoln "CA_LOCAL_VERSION=$CA_LOCAL_VERSION"
    infoln "CA_DOCKER_IMAGE_VERSION=$CA_DOCKER_IMAGE_VERSION"
    if [ "$CA_LOCAL_VERSION" != "$CA_DOCKER_IMAGE_VERSION" ]; then
        warnln "Local fabric-ca binaries and docker images are out of sync. This may cause problems."
    fi
}

function createOrgs() {
    infoln "Generating certificates using Fabric CA"
    ${CONTAINER_CLI_COMPOSE} -f compose/$COMPOSE_FILE_CA up -d 2>&1
    . scripts/registerEnroll.sh

    while :
    do
      if [ ! -f "organizations/fabric-ca/component_supplier/tls-cert.pem" ]; then
        sleep 1
      else
        break
      fi
    done

    infoln "Creating component_supplier Identities"
    createOrgComponentSupplier
    infoln "Creating manufacturer Identities"
    createOrgManufacturer
    infoln "Creating insurer Identities"
    createOrgInsurer
    infoln "Creating store Identities"
    createOrgStore
    infoln "Creating maintenancer Identities"
    createOrgMaintenancer
    infoln "Creating consumer Identities"
    createOrgConsumer

    infoln "Generating CCP files for every organization"
    ./scripts/ccp_generate.sh
}
# . scripts/network.config
prereqsInfo
createOrgs
COMPOSE_FILES="-f compose/${COMPOSE_FILE_CAR}"
DOCKER_SOCK="${DOCKER_SOCK}" ${CONTAINER_CLI_COMPOSE} ${COMPOSE_FILES} up -d 2>&1
$CONTAINER_CLI ps -a

