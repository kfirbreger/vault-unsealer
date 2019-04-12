#!/bin/sh

# Creating a network
docker network create -d bridge safe

# Starting the containers
docker-compose --project-name unsealer up -d --force-recreate
# Initiating vault
docker exec unsealer_vault-0_1 vault operator init | grep "Unseal Key" | awk '{print $4}' > keys.txt

# Starting the unsealer
# This should lead to the vaults being unsealed
docker build ../../ --tag unsealer
docker run --rmi unsealer --network=safe


# Cleaning up
#rm keys.txt

