BLOCKCHAIN_HOME=${BLOCKCHAIN_HOME:-${PWD}}
. ${BLOCKCHAIN_HOME}/scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export Component_Supplier_CA=${BLOCKCHAIN_HOME}/organizations/component_supplier.carlife.com/tlsca/tlsca.component_supplier.carlife.com-cert.pem
export Manufacturer_CA=${BLOCKCHAIN_HOME}/organizations/manufacturer.carlife.com/tlsca/tlsca.manufacturer.carlife.com-cert.pem
export Store_CA=${BLOCKCHAIN_HOME}/organizations/store.carlife.com/tlsca/tlsca.store.carlife.com-cert.pem
export Insurer_CA=${BLOCKCHAIN_HOME}/organizations/insurer.carlife.com/tlsca/tlsca.insurer.carlife.com-cert.pem
export Maintenancer_CA=${BLOCKCHAIN_HOME}/organizations/maintenancer.carlife.com/tlsca/tlsca.maintenancer.carlife.com-cert.pem
export Consumer_CA=${BLOCKCHAIN_HOME}/organizations/consumer.carlife.com/tlsca/tlsca.consumer.carlife.com-cert.pem


setGlobals() {
  local USING_ORG=""
  local USING_NODE=""
  local PORT=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  USING_NODE=$2
  infoln "Using organization ${USING_ORG}"
  if [[ $USING_ORG == "component_supplier" ]]; then
    export CORE_PEER_LOCALMSPID=ComponentSupplierMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$Component_Supplier_CA
    export CORE_PEER_MSPCONFIGPATH=${BLOCKCHAIN_HOME}/organizations/component_supplier.carlife.com/users/Admin1@component_supplier.carlife.com/msp
    if [[ $USING_NODE == "orderer1" ]]; then
      PORT=7050
    elif [[ $USING_NODE == "cartire" ]]; then
      PORT=7051
    elif [[ $USING_NODE == "carbody" ]]; then
      PORT=7052
    elif [[ $USING_NODE == "carinterior" ]]; then
      PORT=7053
    else
      errorln "NODE Unknown"        
    fi
    export CORE_PEER_ADDRESS=localhost:${PORT}
  elif [[ $USING_ORG == "manufacturer" ]]; then
    export CORE_PEER_LOCALMSPID=ManufacturerMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$Manufacturer_CA
    export CORE_PEER_MSPCONFIGPATH=${BLOCKCHAIN_HOME}/organizations/manufacturer.carlife.com/users/Admin1@manufacturer.carlife.com/msp
    if [[ $USING_NODE == "orderer2" ]]; then
      PORT=7054
    elif [[ $USING_NODE == "factory1" ]]; then
      PORT=7055
    elif [[ $USING_NODE == "factory2" ]]; then
      PORT=7056
    else
      errorln "NODE Unknown"        
    fi  
    export CORE_PEER_ADDRESS=localhost:${PORT}
  elif [[ $USING_ORG == "store" ]]; then
    export CORE_PEER_LOCALMSPID=StoreMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$Store_CA
    export CORE_PEER_MSPCONFIGPATH=${BLOCKCHAIN_HOME}/organizations/store.carlife.com/users/Admin1@store.carlife.com/msp
    if [[ $USING_NODE == "sailer1" ]]; then
      PORT=7057
    elif [[ $USING_NODE == "sailer2" ]]; then
      PORT=7058
    else
      errorln "NODE Unknown"        
    fi     
    export CORE_PEER_ADDRESS=localhost:${PORT}
  elif [[ $USING_ORG == "insurer" ]]; then
    export CORE_PEER_LOCALMSPID=InsurerMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$Insurer_CA
    export CORE_PEER_MSPCONFIGPATH=${BLOCKCHAIN_HOME}/organizations/insurer.carlife.com/users/Admin1@insurer.carlife.com/msp
    if [[ $USING_NODE == "orderer3" ]]; then
      PORT=7059
    elif [[ $USING_NODE == "pcompany" ]]; then
      PORT=7060
    elif [[ $USING_NODE == "rcompany" ]]; then
      PORT=7061
    else
      errorln "NODE Unknown"        
    fi      
    export CORE_PEER_ADDRESS=localhost:${PORT}
  elif [[ $USING_ORG == "maintenancer" ]]; then
    export CORE_PEER_LOCALMSPID=MaintenancerMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$Maintenancer_CA
    export CORE_PEER_MSPCONFIGPATH=${BLOCKCHAIN_HOME}/organizations/maintenancer.carlife.com/users/Admin1@maintenancer.carlife.com/msp
    if [[ $USING_NODE == "fixer1" ]]; then
      PORT=7062
    elif [[ $USING_NODE == "fixer2" ]]; then
      PORT=7063
    else
      errorln "NODE Unknown"        
    fi      
    export CORE_PEER_ADDRESS=localhost:${PORT}
  elif [[ $USING_ORG == "consumer" ]]; then
    export CORE_PEER_LOCALMSPID=ConsumerMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$Consumer_CA
    export CORE_PEER_MSPCONFIGPATH=${BLOCKCHAIN_HOME}/organizations/consumer.carlife.com/users/Admin1@consumer.carlife.com/msp
    if [[ $USING_NODE == "user1" ]]; then
      PORT=7064
    else
      errorln "NODE Unknown"
    fi
    export CORE_PEER_ADDRESS=localhost:${PORT}     
  else
    errorln "ORG Unknown"
  fi

  if [ "$VERBOSE" = "true" ]; then
    env | grep CORE
  fi
}

parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""
  while [ "$#" -gt 0 ]; do
    setGlobals $1 $2
    PEER="$2.$1"
    ## Set peer addresses
    if [ -z "$PEERS" ]
    then
	PEERS="$PEER"
    else
	PEERS="$PEERS $PEER"
    fi
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
    ## Set path to TLS certificate
    CA=$CORE_PEER_TLS_ROOTCERT_FILE
    TLSINFO=(--tlsRootCertFiles "${!CA}")
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
    # shift by one to get to the next organization
    shift 2
  done
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$2"
  fi
}