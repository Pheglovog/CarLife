#!/bin/bash

source scripts/utils.sh

CHANNEL_NAME="carchannel"
CC_NAME="carlife-chaincode-go"
CC_SRC_PATH="./chaincode-car"
CC_SRC_LANGUAGE="go"
CC_VERSION="1.0"
CC_SEQUENCE="1"
DELAY="3"
MAX_RETRY="5"
VERBOSE="false"

println "executing with the following"
println "- CHANNEL_NAME: ${C_GREEN}${CHANNEL_NAME}${C_RESET}"
println "- CC_NAME: ${C_GREEN}${CC_NAME}${C_RESET}"
println "- CC_SRC_PATH: ${C_GREEN}${CC_SRC_PATH}${C_RESET}"
println "- CC_SRC_LANGUAGE: ${C_GREEN}${CC_SRC_LANGUAGE}${C_RESET}"
println "- CC_VERSION: ${C_GREEN}${CC_VERSION}${C_RESET}"
println "- CC_SEQUENCE: ${C_GREEN}${CC_SEQUENCE}${C_RESET}"
println "- DELAY: ${C_GREEN}${DELAY}${C_RESET}"
println "- MAX_RETRY: ${C_GREEN}${MAX_RETRY}${C_RESET}"
println "- VERBOSE: ${C_GREEN}${VERBOSE}${C_RESET}"


export FABRIC_CFG_PATH=${PWD}/config

. scripts/envVar.sh
. scripts/ccutils.sh

INIT_REQUIRED=""
CC_END_POLICY=""
CC_COLL_CONFIG=""


function checkPrereqs() {
  jq --version > /dev/null 2>&1

  if [[ $? -ne 0 ]]; then
    errorln "jq command not found..."
    exit 1
  fi
}

checkPrereqs

./scripts/packageCC.sh $CC_NAME $CC_SRC_PATH $CC_SRC_LANGUAGE $CC_VERSION

PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid ${CC_NAME}.tar.gz)

infoln "Installing chaincode on component_supplier peers..."
installChaincode component_supplier cartire
installChaincode component_supplier carbody
installChaincode component_supplier carinterior
infoln "Installing chaincode on manufacturer peers..."
installChaincode manufacturer factory1
installChaincode manufacturer factory2
infoln "Installing chaincode on store peers..."
installChaincode store sailer1
installChaincode store sailer2
infoln "Installing chaincode on insurer peers..."
installChaincode insurer pcompany
installChaincode insurer rcompany
infoln "Installing chaincode on maintenancer peers..."
installChaincode maintenancer fixer1
installChaincode maintenancer fixer2
infoln "Installing chaincode on consumer peers..."
installChaincode consumer user1

resolveSequence

queryInstalled component_supplier cartire
approveForMyOrg component_supplier cartire
checkCommitReadiness component_supplier cartire "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": false" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness manufacturer factory1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": false" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness store sailer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": false" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness insurer pcompany "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": false" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness maintenancer fixer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": false" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness consumer user1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": false" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"



queryInstalled manufacturer factory1
approveForMyOrg manufacturer factory1
checkCommitReadiness component_supplier cartire "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness manufacturer factory1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness store sailer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness insurer pcompany "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness maintenancer fixer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness consumer user1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": false" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"


queryInstalled store sailer1
approveForMyOrg store sailer1
checkCommitReadiness component_supplier cartire "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness manufacturer factory1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness store sailer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness insurer pcompany "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness maintenancer fixer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness consumer user1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": false" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"


queryInstalled insurer pcompany
approveForMyOrg insurer pcompany
checkCommitReadiness component_supplier cartire "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness manufacturer factory1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness store sailer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness insurer pcompany "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness maintenancer fixer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"
checkCommitReadiness consumer user1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": false" "\"ConsumerMSP\": false"


queryInstalled maintenancer fixer1
approveForMyOrg maintenancer fixer1
checkCommitReadiness component_supplier cartire "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": false"
checkCommitReadiness manufacturer factory1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": false"
checkCommitReadiness store sailer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": false"
checkCommitReadiness insurer pcompany "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": false"
checkCommitReadiness maintenancer fixer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": false"
checkCommitReadiness consumer user1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": false"


queryInstalled consumer user1
approveForMyOrg consumer user1
checkCommitReadiness component_supplier cartire "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": true"
checkCommitReadiness manufacturer factory1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": true"
checkCommitReadiness store sailer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": true"
checkCommitReadiness insurer pcompany "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": true"
checkCommitReadiness maintenancer fixer1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": true"
checkCommitReadiness consumer user1 "\"ComponentSupplierMSP\": true" "\"ManufacturerMSP\": true" "\"StoreMSP\": true" "\"InsurerMSP\": true" "\"MaintenancerMSP\": true" "\"ConsumerMSP\": true"

commitChaincodeDefinition component_supplier cartire manufacturer factory1 store sailer1 insurer pcompany maintenancer fixer1 consumer user1

queryCommitted component_supplier cartire 
queryCommitted manufacturer factory1 
queryCommitted store sailer1 
queryCommitted insurer pcompany 
queryCommitted maintenancer fixer1 
queryCommitted consumer user1 

exit 0

