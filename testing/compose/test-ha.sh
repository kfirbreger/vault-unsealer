#!/bin/sh

# Creating a network
docker network create -d bridge safe
sleep 2
# Starting the containers
docker-compose --project-name unsealer up -d --force-recreate 
# Giving vault time to start
echo "Waiting 10 second for vault to start up"
sleep 10
# Initiating vault
docker exec unsealer_vault-0_1 vault operator init | grep "Unseal Key" | awk '{print $4}' > keys.txt

# Starting the unsealer
# This should lead to the vaults being unsealed
docker build -f Dockerfile ../../ --tag unsealer
docker run --rm --network=safe --cap-add=IPC_LOCK unsealer

# Cleaning up
#rm keys.txt

