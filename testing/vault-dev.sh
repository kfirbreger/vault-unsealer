#!/bin/sh
docker run --cap-add=IPC_LOCK --rm -d --name=vault-dev -e 'VAULT_ADDR=http://127.0.0.1:8200' vault
