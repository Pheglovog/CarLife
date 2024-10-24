#!/bin/bash
ROOTDIR=$(cd "$(dirname "$0")" && pwd)
export PATH=${ROOTDIR}/bin:${PWD}/bin:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx
export VERBOSE=false

. scripts/utils.sh
infoln "Creating channel carchannel."

if ! $CONTAINER_CLI info > /dev/null 2>&1 ; then
    fatalln "$CONTAINER_CLI network is required to be running to create a channel"
fi

. scripts/envVar.sh

CHANNEL_NAME=carchannel
DELAY=3
MAX_RETRY=5
VERBOSE=false
CONTAINER_CLI=docker
CONTAINER_CLI_COMPOSE=docker-compose
infoln "Using ${CONTAINER_CLI} and ${CONTAINER_CLI_COMPOSE}"

if [ ! -d "channel-artifacts" ]; then
	mkdir channel-artifacts
fi

createChannelGenesisBlock() {
    setGlobals component_supplier cartire
    which configtxgen
    if [ "$?" -ne 0 ]; then
        fatalln "configtxgen tool not found."
    fi
    set -x
    configtxgen -profile ChannelUsingRaft -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
    res=$?
    { set +x; } 2>/dev/null
    verifyResult $res "Failed to generate channel configuration transaction..."
}

createChannel() {
	# Poll in case the raft leader is not set yet
	local rc=1
	local COUNTER=1
	infoln "Adding orderers"
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
		sleep $DELAY
		set -x
        . scripts/orderer.sh ${CHANNEL_NAME}> /dev/null 2>&1
		res=$?
		{ set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "Channel creation failed"
}

joinChannel() {
    ORG=$1
    NODE=$2
    FABRIC_CFG_PATH=$PWD/config/
    setGlobals $ORG $NODE
    local rc=1
    local COUNTER=1
    ## Sometimes Join takes time, hence retry
    while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
        sleep $DELAY
        set -x
        peer channel join -b $BLOCKFILE >&log.txt
        res=$?
        { set +x; } 2>/dev/null
        let rc=$res
        COUNTER=$(expr $COUNTER + 1)
    done
    cat log.txt
    verifyResult $res "After $MAX_RETRY attempts, peer0.org${ORG} has failed to join channel '$CHANNEL_NAME' "
}

setAnchorPeer() {
    ORG=$1
    NODE=$2
    . scripts/setAnchorPeer.sh $ORG $NODE $CHANNEL_NAME 
}

## Create channel genesis block
FABRIC_CFG_PATH=${PWD}/config/
BLOCKFILE="./channel-artifacts/${CHANNEL_NAME}.block"

infoln "Generating channel genesis block '${CHANNEL_NAME}.block'"
FABRIC_CFG_PATH=${PWD}/configtx
createChannelGenesisBlock

## Create channel
infoln "Creating channel ${CHANNEL_NAME}"
createChannel
successln "Channel '$CHANNEL_NAME' created"

## Join all the peers to the channel
infoln "Joining component_supplier peers to the channel..."
joinChannel component_supplier cartire
joinChannel component_supplier carbody
joinChannel component_supplier carinterior
infoln "Joining manufacturer peers to the channel..."
joinChannel manufacturer factory1
joinChannel manufacturer factory2
infoln "Joining store peers to the channel..."
joinChannel store sailer1
joinChannel store sailer2
infoln "Joining insurer peers to the channel..."
joinChannel insurer pcompany
joinChannel insurer rcompany
infoln "Joining maintenancer peers to the channel..."
joinChannel maintenancer fixer1
joinChannel maintenancer fixer2
infoln "Joining consumer peers to the channel..."
joinChannel consumer user1

## Set the anchor peers for each org in the channel
infoln "Setting anchor peer for component_supplier..."
setAnchorPeer component_supplier cartire
infoln "Setting anchor peer for manufacturer..."
setAnchorPeer manufacturer factory1
infoln "Setting anchor peer for store..."
setAnchorPeer store sailer1
infoln "Setting anchor peer for insurer..."
setAnchorPeer insurer pcompany
infoln "Setting anchor peer for maintenancer..."
setAnchorPeer maintenancer fixer1
infoln "Setting anchor peer for consumer..."
setAnchorPeer consumer user1

successln "Channel '$CHANNEL_NAME' joined"