#!/bin/sh
docker run --rm --name vault -d -p 8200:8200 --cap-add=IPC_LOCK -e 'VAULT_LOCAL_CONFIG={"backend": {"file": {"path": "/vault/file"}}, "listener": { "tcp": {"address": "0.0.0.0:8200", "tls_disable": "1"}}}' -e 'VAULT_ADDR=http://127.0.0.1:8200' vault server
