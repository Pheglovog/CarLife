#!/bin/bash

########################orderer1######################
channel_name=$1
export ORDERER_PORT=8050
export ORDERER_CA=$Component_Supplier_CA
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/component_supplier.carlife.com/orderers/orderer1.component_supplier.carlife.com/tls/server.crt /dev/null 2>&1
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/component_supplier.carlife.com/orderers/orderer1.component_supplier.carlife.com/tls/server.key /dev/null 2>&1

osnadmin channel join --channelID ${channel_name} --config-block ./channel-artifacts/${channel_name}.block -o localhost:${ORDERER_PORT} --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY" >> log.txt 2>&1

#######################orderer2######################
export ORDERER_PORT=8054
export ORDERER_CA=$Manufacturer_CA
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/manufacturer.carlife.com/orderers/orderer2.manufacturer.carlife.com/tls/server.crt /dev/null 2>&1
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/manufacturer.carlife.com/orderers/orderer2.manufacturer.carlife.com/tls/server.key /dev/null 2>&1

osnadmin channel join --channelID ${channel_name} --config-block ./channel-artifacts/${channel_name}.block -o localhost:${ORDERER_PORT} --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY" >> log.txt 2>&1

#######################orderer3######################
export ORDERER_PORT=8059
export ORDERER_CA=$Insurer_CA
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/insurer.carlife.com/orderers/orderer3.insurer.carlife.com/tls/server.crt /dev/null 2>&1
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/insurer.carlife.com/orderers/orderer3.insurer.carlife.com/tls/server.key /dev/null 2>&1

osnadmin channel join --channelID ${channel_name} --config-block ./channel-artifacts/${channel_name}.block -o localhost:${ORDERER_PORT} --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY" >> log.txt 2>&1
