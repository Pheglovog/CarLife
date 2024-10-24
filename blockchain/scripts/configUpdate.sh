#!/bin/bash

BLOCKCHAIN_HOME=${BLOCKCHAIN_HOME:-${PWD}}
. ${BLOCKCHAIN_HOME}/scripts/envVar.sh

# fetchChannelConfig <org> <channel_id> <output_json>
# Writes the current channel config for a given channel to a JSON file
# NOTE: this requires jq and configtxlator for execution.
fetchChannelConfig() {
  ORG=$1
  NODE=$2
  CHANNEL=$3
  OUTPUT=$4

  setGlobals $ORG $NODE

  infoln "Fetching the most recent configuration block for the channel"
  set -x
  peer channel fetch config ${BLOCKCHAIN_HOME}/channel-artifacts/config_block.pb -o localhost:7050 --ordererTLSHostnameOverride orderer1.component_supplier.carlife.com -c $CHANNEL --tls --cafile "$Component_Supplier_CA"
  { set +x; } 2>/dev/null

  infoln "Decoding config block to JSON and isolating config to ${OUTPUT}"
  set -x
  configtxlator proto_decode --input ${BLOCKCHAIN_HOME}/channel-artifacts/config_block.pb --type common.Block --output ${BLOCKCHAIN_HOME}/channel-artifacts/config_block.json
  jq .data.data[0].payload.data.config ${BLOCKCHAIN_HOME}/channel-artifacts/config_block.json >"${OUTPUT}"
  res=$?
  { set +x; } 2>/dev/null
  verifyResult $res "Failed to parse channel configuration, make sure you have jq installed"
}

# createConfigUpdate <channel_id> <original_config.json> <modified_config.json> <output.pb>
# Takes an original and modified config, and produces the config update tx
# which transitions between the two
# NOTE: this requires jq and configtxlator for execution.
createConfigUpdate() {
  CHANNEL=$1
  ORIGINAL=$2
  MODIFIED=$3
  OUTPUT=$4

  set -x
  configtxlator proto_encode --input "${ORIGINAL}" --type common.Config --output ${BLOCKCHAIN_HOME}/channel-artifacts/original_config.pb
  configtxlator proto_encode --input "${MODIFIED}" --type common.Config --output ${BLOCKCHAIN_HOME}/channel-artifacts/modified_config.pb
  configtxlator compute_update --channel_id "${CHANNEL}" --original ${BLOCKCHAIN_HOME}/channel-artifacts/original_config.pb --updated ${BLOCKCHAIN_HOME}/channel-artifacts/modified_config.pb --output ${BLOCKCHAIN_HOME}/channel-artifacts/config_update.pb
  configtxlator proto_decode --input ${BLOCKCHAIN_HOME}/channel-artifacts/config_update.pb --type common.ConfigUpdate --output ${BLOCKCHAIN_HOME}/channel-artifacts/config_update.json
  echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL'", "type":2}},"data":{"config_update":'$(cat ${BLOCKCHAIN_HOME}/channel-artifacts/config_update.json)'}}}' | jq . > ${BLOCKCHAIN_HOME}/channel-artifacts/config_update_in_envelope.json
  configtxlator proto_encode --input ${BLOCKCHAIN_HOME}/channel-artifacts/config_update_in_envelope.json --type common.Envelope --output "${OUTPUT}"
  { set +x; } 2>/dev/null
}

# signConfigtxAsPeerOrg <org> <configtx.pb>
# Set the peerOrg admin of an org and sign the config update
signConfigtxAsPeerOrg() {
  ORG=$1
  NODE=$2
  CONFIGTXFILE=$3
  setGlobals $ORG $NODE
  set -x
  peer channel signconfigtx -f "${CONFIGTXFILE}"
  { set +x; } 2>/dev/null
}
