#!/bin/bash

docker-compose -f compose/docker-compose-car.yaml down -v
docker-compose -f compose/docker-compose-ca.yaml down -v

# Specify the target directory
TARGET_DIR="organizations"

# Check if the target directory exists
if [ ! -d "$TARGET_DIR" ]; then
  echo "Directory $TARGET_DIR does not exist!"
  exit 1
fi
rm -rf $TARGET_DIR
rm -rf channel-artifacts
echo "Deletion operation completed."