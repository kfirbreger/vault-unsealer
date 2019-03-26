package internal

import (
	"flag"
	"fmt"
    "log"
    "os"
    "os/user"
    "path/filepath"

	"github.com/BurntSushi/toml"
)

type Instance struct {
	Domain string
}

type VaultConf struct {
	Protocol       string
	StatusPath     string `toml:"status_path"`
	UnsealPath     string `toml:"unseal_path"`
	UnsealKeyCount int    `toml:"unseal_key_count"`
	CheckInterval  int    `toml:"seal_check_interval"`
}

type WorkersConf struct {
	StatusCheckCount int `toml:"status_check"`
	UnsealCount      int `toml:"unseal"`
	LoggingCount     int `toml:"logging"`
}

type Service struct {
	Vault   VaultConf
	Workers WorkersConf
	Servers []Instance `toml:"server"`
}

func Load(filepath string, s *Service) {
	if _, err := toml.DecodeFile(filepath, s); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*s)
	// @todo add verification
}

type CliParams struct {
    ConfigFile     string
	UnsealKeyCount int
	StatusPath     string
	UnsealPath     string
	Interval       int
	Protocol       string
}

// Handling the parameters
func getCliParams() CliParams {
	var params CliParams

    params.ConfigFile = *flag.String("config", "", "Specify the configuration file to use")
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

func LoadConfiguration() *Service {
	var conf Service
    cliParams := getCliParams()
    // Reading the config file
    configFile := "./config.toml"
    if len(cliParams.ConfigFile) > 0 {
        if _, err := os.Stat(cliParams.ConfigFile); err == nil {
            // There is a file with that name, so use it as a config file
            // It is safe to reuse err as it is nil
            if configFile, err = expand(cliParams.ConfigFile); err != nil {
                log.Fatal(err)
            }
        } else if os.IsNotExist(err) {
            log.Fatalf("The file %s does not seem to exist", cliParams.ConfigFile)
        } else {
            // Something is wrong, display the error and terminate
            log.Fatal(err)
        }
    }

	Load(configFile, &conf)
	// TODO add cli params
	return &conf
}

func expand(path string) (string, error) {
    /**
    Expand the directory to be absolute instead of relative path
    as using ~ will not always work when opening files
    */
    if len(path) == 0 || path[0] != '~' {
        return path, nil
    }

    usr, err := user.Current()
    if err != nil {
        return "", err
    }
    return filepath.Join(usr.HomeDir, path[1:]), nil
}
