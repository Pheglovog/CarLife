#!/bin/bash

# to .. as relative path to make the import works
BLOCKCHAIN_HOME=${BLOCKCHAIN_HOME:-${PWD}}
. ${BLOCKCHAIN_HOME}/scripts/configUpdate.sh


# NOTE: This requires jq and configtxlator for execution.
createAnchorPeerUpdate() {
  infoln "Fetching channel config for channel $CHANNEL_NAME"
  fetchChannelConfig $ORG $NODE $CHANNEL_NAME ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}config.json

  infoln "Generating anchor peer update transaction for ${ORG} on channel $CHANNEL_NAME"

  if [[ $ORG == "component_supplier" ]]; then
    HOST="cartire.component_supplier.carlife.com"
    PORT=7051
  elif [[ $ORG == "manufacturer" ]]; then
    HOST="factory1.manufacturer.carlife.com"
    PORT=7055
  elif [[ $ORG == "store" ]]; then
    HOST="sailer1.store.carlife.com"
    PORT=7057
  elif [[ $ORG == "insurer" ]]; then
    HOST="pcompany.insurer.carlife.com"
    PORT=7060     
  elif [[ $ORG == "maintenancer" ]]; then
    HOST="fixer1.maintenancer.carlife.com"
    PORT=7062
  elif [[ $ORG == "consumer" ]]; then
    HOST="user1.consumer.carlife.com"
    PORT=7064    
  else
    errorln "ORG ${ORG} unknown"
  fi

  set -x
  # Modify the configuration to append the anchor peer 
  jq '.channel_group.groups.Application.groups.'${CORE_PEER_LOCALMSPID}'.values += {"AnchorPeers":{"mod_policy": "Admins","value":{"anchor_peers": [{"host": "'$HOST'","port": '$PORT'}]},"version": "0"}}' ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}config.json > ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}modified_config.json
  res=$?
  { set +x; } 2>/dev/null
  verifyResult $res "Channel configuration update for anchor peer failed, make sure you have jq installed"
  

  # Compute a config update, based on the differences between 
  # {orgmsp}config.json and {orgmsp}modified_config.json, write
  # it as a transaction to {orgmsp}anchors.tx
  createConfigUpdate ${CHANNEL_NAME} ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}config.json ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}modified_config.json ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx
}

signAnchorPeerUpdate() {
  TX=${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx
  signConfigtxAsPeerOrg component_supplier cartire $TX
  signConfigtxAsPeerOrg manufacturer factory1 $TX  
  signConfigtxAsPeerOrg store sailer1 $TX  
  signConfigtxAsPeerOrg insurer pcompany $TX  
  signConfigtxAsPeerOrg maintenancer fixer1 $TX 
}
updateAnchorPeer() {  
  set -x
  peer channel update -o localhost:7050 --ordererTLSHostnameOverride orderer1.component_supplier.carlife.com -c $CHANNEL_NAME -f ${BLOCKCHAIN_HOME}/channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls --cafile "$Component_Supplier_CA" >&log.txt
  { set +x; } 2>/dev/null
  res=$?
  cat log.txt
  verifyResult $res "Anchor peer update failed"
  successln "Anchor peer set for org '$CORE_PEER_LOCALMSPID' on channel '$CHANNEL_NAME'"
}

ORG=$1
NODE=$2
CHANNEL_NAME=$3

setGlobals $ORG $NODE
createAnchorPeerUpdate 

signAnchorPeerUpdate

setGlobals $1 $2
updateAnchorPeer 
