package internal

import (
    "flag"
	"github.com/BurntSushi/toml"
)

type Instance struct {
	Domain string
}

type ServersList struct {
	Instances []Instance
}

type VaultConf struct {
	Protocol       string
	StatusPath     string `toml:status_path`
	UnsealPath     string `toml:unseal_path`
	UnsealKeyCount string `toml:unseal_key_path`
	CheckInterval  int    `toml:seal_check_interval`
}

type WorkersConf struct {
	StatusCheckCount int `toml:status_check`
	UnsealCount      int `toml:unseal`
	LoggingCount     int `toml:logging`
}

type Service struct {
	Vault   VaultConf
	Workers WorkersConf
	Servers ServersList
}

func (s *Service) Load(filepath string) error {
	_, err := toml.DecodeFile(filepath, s)
	// @todo add verification
	return err
}

type CliParams struct {
	UnsealKeyCount int
	StatusPath     string
	UnsealPath     string
	Interval       int
	Protocol       string
}

// Handling the parameters
func getCliParams() CliParams {
	var params CliParams

	params.UnsealKeyCount = *flag.Int("unsealing-keys", 0, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
	params.StatusPath = *flag.String("status-path", "", "Give a custom status check path")
	params.UnsealPath = *flag.String("unseal-path", "", "Give a custom unseal path")
	params.Interval = *flag.Int("check-interval", 0, "The status check interval")
	params.Protocol = *flag.String("protocol", "", "Use a custom protocol")

	flag.Parse()

	return params
}

// Updating config from CLI, if needed
func updateConfig(serv *Service, params *CliParams) {
	return
}
