# Vault unsealer

Providing vault auto-unsealing when cloud infra is not available

## Layout

Project layout follows the guidelines in the golang standards [prjocet layout](https://github.com/golang-standards/project-layout)

### Libraries
Vault unsealer tries to keep external depencies to a minimum, for security reasons. Currently, these are the depencies for the project:

- [Memguard](https://github.com/awnumar/memguard) used to store the unsealing keys safely in memory.
- [toml](https://github.com/BurntSushi/toml) used to read and parse the toml configuration file

