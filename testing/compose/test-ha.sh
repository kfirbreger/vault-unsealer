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
docker run --name=unsealer -d --rm --network=safe --cap-add=IPC_LOCK unsealer

echo "Giving everything a minute to stabilize"
sleep 60

echo "Showing unsealer logs:"
echo "======================"
docker logs unsealer

echo "Showing vault-0 status"
docker-compose --project-name unsealer exec vault-0 vault status

echo "Showing vault-1 status"
docker-compose --project-name unsealer exec vault-1 vault status

echo "Restarting a vault server"
docker-compose --project-name unsealer restart vault-1

echo "Giving it 10 seconds"
sleep 10

echo "Showing vault-1 status"
docker-compose --project-name unsealer exec vault-1 vault status

echo "Cleaning it all up"

./cleanup.sh
