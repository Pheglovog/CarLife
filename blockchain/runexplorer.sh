#!/bin/bash
export EXPLORER_CONFIG_FILE_PATH=./config.json
export EXPLORER_PROFILE_DIR_PATH=./connection-profile
export FABRIC_CRYPTO_PATH=./organizations
export PORT=3030
cp -r organizations ./explorer/
chown -R $USER organizations
mv ./explorer/organizations/store.carlife.com/users/Admin1@store.carlife.com/msp/keystore/*sk ./explorer/organizations/store.carlife.com/users/Admin1@store.carlife.com/msp/keystore/key.pem
docker-compose -f explorer/docker-compose.yaml up -d
