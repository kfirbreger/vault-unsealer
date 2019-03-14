package unsealer

import (
    "github.com/BurntSushi/toml"
)

type Instance struct {
    Domain string
}

type ServersList struct {
    Instances []Instance
}

type VaultConf struct {
    Protocol string
    StatusPath string `toml:status_path`
    UnsealPath string `toml:unseal_path`
    UnsealKeyCount string `toml:unseal_key_path`
    CheckInterval int `toml:seal_check_interval`
}

type WorkersConf struct {
    StatusCheckCount int `toml:status_check`
    UnsealCount int `toml:unseal`
    LoggingCount int `toml:logging`
}

type Service struct {
    Vault VaultConf
    Workers WorkersConf
    Servers ServersList
}

func (*s Service) load(filepath string) error {
    var err error
    _, err = toml.decodefile(filepath, s)
    // @todo add verification
    return err
}


type CliParams struct {
    UnsealKeyCount int
    InstanceDomain []string
    InstanceReset bool
    StatusPath string
    UnsealPath string
    Interval int
    Protocol string
}

// Handling the parameters
func getCliParams() CliParams {
    var params CliParams

    params.UnsealKeyCount = flag.Int("unsealing-keys", nil, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
    params.InstanceDomain = flag.String("instance", "", "Add a Vault instance to the monitoring list")
    params.InstanceReset = flag.Bool("instance-reset", false, "Remove all instances from the config")
    params.StatusPath = flag.String("status-path", "", "Give a custom status check path")
    params.UnsealPath = flag.String("unseal-path", "", "Give a custom unseal path")
    params.Interval = flag.Int("check-interval", nil, "The status check interval")
    params.Protocol = flag.String("protocol", "", "Use a custom protocol")

    flag.Parse()

    return params
}

// Updating config from CLI, if needed
func updateConfig(serv *Service, params *CliParams) {
    return nil
}

