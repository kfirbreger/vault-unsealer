#!/bin/sh
docker run --cap-add=IPC_LOCK --rm -d --name=vault-dev -p 8200:8200 -e 'VAULT_ADDR=http://0.0.0.0:8200' vault
