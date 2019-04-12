#!/bin/sh

# Creating a network
docker network create safe

# Starting the containers
docker-compose up -d --project-name unsealer

# Initiating vault
docker-compose exec vault-0 vault operator init | grep "Unseal Key" | awk '{print $4}' > keys.txt

# Starting the unsealer
# This should lead to the vaults being unsealed
docker build . --name unsealer
docker run --rmi unsealer


# Cleaning up
#rm keys.txt

