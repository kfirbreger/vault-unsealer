package unsealer

import (
    "github.com/BurntSushi/toml"
)

type instance struct {
    Domain string
}

type servers struct {
    Instance []instance
}

type vault struct {
    Protocol string
    StatusPath string `toml:status_path`
    UnsealPath string `toml:unseal_path`
    UnsealKeyCount string `toml:unseal_key_path`
    CheckInterval int `toml:seal_check_interval`
}

type service struct {
    Vault vault
    Servers servers
}

func (*s service) Load(filePath string) error {
    var err error
    _, err = toml.DecodeFile(filePath, s)
    // @TODO add verification
    return err
}
