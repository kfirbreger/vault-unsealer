# Vault unsealer

Providing vault auto-unsealing when cloud infra is not available

## Installation

Releases, including binaries are available on [Github](https://github.com/kfirbreger/vault-unsealer/releases). It is recommended to use them.
The releases include binaries for Linux and macOS. The other option is to build it yourself

    git clone https://github.com/kfirbreger/vault-unsealer
    go build -o unsealer cmd/unsealer/main.go

As per usual the code needs to in your `GOPATH`.

## Usage

The vault unsealer uses a toml config file for initialization. An example of the configuration file can be found in the __configs__ directory.
By default the unsealer expects the configuration file to be called config.toml and be present in the same directory as the binary. This behavior
can be overwritten by passing the `-config' flag followed by the path to the config. To get a full list of the supported flags run

    unsealer --help

Command line flags will override values given in the config file. Even though it is possible to pass all the needed values via command line flags,
currently a config file is still required.

Passing the keys via a file or via a cli flag, though supported is only meant for development. __Do not used this in production!__

## Development

If you wish to contribute to the project here are some tips

### Layout

Project layout follows the guidelines in the golang standards [project layout](https://github.com/golang-standards/project-layout)

### Libraries
Vault unsealer tries to keep external depencies to a minimum, for security reasons. Currently, these are the depencies for the project:

- [Memguard](https://github.com/awnumar/memguard) used to store the unsealing keys safely in memory.
- [toml](https://github.com/BurntSushi/toml) used to read and parse the toml configuration file

