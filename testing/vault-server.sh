#!/bin/sh
docker run --rm --name vault -d -p 8200:8200 --cap-add=IPC_LOCK -e 'VAULT_LOCAL_CONFIG={"backend": {"file": {"path": "/vault/file"}}, "listener": { "tcp": {"address": "0.0.0.0:8200", "tls_disable": "1"}}}' -e 'VAULT_ADDR=http://127.0.0.1:8200' vault server

# Giving docker 2 seconds to start it all up
sleep 2

# Init the server. Get the unseal keys only and save them to a text file
docker exec vault vault operator init | grep "Unseal Key" | awk '{print $4}' > keys.txt


