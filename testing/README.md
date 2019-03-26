# Testing helper functions

In order to properly test the unsealer, it is needed to be able to run vault. The scripts here are used to start vault in a docker container.

### Requirements
Docker needs to be installed for the scrupts to work. The scripts are written for Unix / Linux systems and will probably not run on windows systems

### Vault in dev mode
Startig the vault in dev mode is a good way to verify that the status check works correctly for an unsealed vault. Since the vault is already unsealed it does not matter what is being filled in the unsealing key section

### Vault in server mode
To test that unsealing is recognized and executed, vault is started as a server. The output of the vault initalization is piped through some commads ending with the unsealing keys being written to a file `keys.txt` This file is than
used by the test framework to load the keys into the unsealer so that unsealing can be tested.
