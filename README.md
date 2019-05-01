[![Go report card](https://goreportcard.com/badge/github.com/kfirbreger/vault-unsealer)](https://goreportcard.com/report/github.com/kfirbreger/vault-unsealer)
# Vault unsealer
Since version 1.0, Vault offers auto unsealing functionality out of the box. For that it uses a public cloud infrastrcture(such as AWS, Azure and Gcloud), to securly store the vault master key.
if you can leverage those services, that should definitely be the way to go. However, there are times in which such services cannot be used(closed networks i.e.). On such systems
Vault currently offers no tools to support automatic unsealing. This project is meant to support auto unsealing while being as secure as possible.

## Target Architecture

The unsealer has been developed with kubernetes in mind. That does not mean that it cannot work on other architectures, just that the focus has been on running the unsealer
as pods on kubernetes. Further, there is the intention to provide support for kubernetes API integration so that if Vault is also running on kubernetes (which Hashicorp currently
[discourages](https://learn.hashicorp.com/vault/operations/production-hardening.html)), the unsealer will listen to the kubernetes API to determine if a vault pod has been
restarted.

## Installation

Releases, including binaries are available on [Github](https://github.com/kfirbreger/vault-unsealer/releases). It is recommended to use them.
The releases include binaries for Linux and macOS. The other option is to build it yourself

    git clone https://github.com/kfirbreger/vault-unsealer
    go build -o unsealer cmd/unsealer/main.go

As per usual the code needs to in your `GOPATH`.

## Usage

The vault unsealer uses a toml config file for initialisation. An example of the configuration file can be found in the __configs__ directory.
By default the unsealer expects the configuration file to be called config.toml and be present in the same directory as the binary. This behaviour
can be overwritten by passing the `-config' flag followed by the path to the config. To get a full list of the supported flags run

    unsealer --help

Command line flags will override values given in the config file. Even though it is possible to pass all the needed values via command line flags,
currently a config file is still required.

Passing the keys via a file or via a cli flag, though supported is only meant for development. __Do not used this in production!__


## Deployment

When moving to production it is recommended to have several unsealer containers running each with insufficient unsealing keys. If one
container is compromised, there is still not enough information to do a seal / unseal operation on the vault. Also it can give robustness
so that even if one unsealer container fails, there are enough running with sufficient unsealing keys to still support auto unsealing.

## Development

Future improvements plans:
- Support the kubernetes API to see if a vault pod has been restarted instead of using http(s) calls
- Using a different vault for keeping the unsealing keys.
- Good logging and monitoring options

If you wish to contribute to the project here are some tips

### Layout

Project layout follows the guidelines in the golang standards [project layout](https://github.com/golang-standards/project-layout)

### Libraries
Vault unsealer tries to keep external decencies to a minimum, for security reasons. Currently, these are the decencies for the project:

- [Memguard](https://github.com/awnumar/memguard) used to store the unsealing keys safely in memory.
- [toml](https://github.com/BurntSushi/toml) used to read and parse the toml configuration file

