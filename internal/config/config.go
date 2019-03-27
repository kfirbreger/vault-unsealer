package config

import (
    "log"
    "os"
    "strings"
)

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
    // Update with the CLI params
    updateConfig(conf, cliParams)

	return &conf
}

// Updating config from CLI, if needed
func updateConfig(conf Service, params CliParams) Service {
    // Checking the flags
    if params.UnsealKeyCount > 0 {
        conf.Vault.UnsealKeyCount = params.UnsealKeyCount
    }

    if params.Interval > 0 {
        conf.Vault.CheckInterval = params.Interval
    }
   
    // converting protocol to lowercase
    prtcl := strings.ToLower(params.Protocol)
    // and checking if its supported
    if prtcl == "http" || prtcl == "https" {
        conf.Vault.Protocol = params.Protocol
    } else if len(params.Protocol) > 0 {
        log.Printf("Unsupported protocol %s given on protocol flag. Using the one defined in the configuration file instead\n", params.Protocol)
    }

    // Checking that if reset instance is givem, at least one instance is also given
    if params.ResetInstances {
        if len(params.Instances) == 0 {
            log.Fatal("An instance reset was passed with no new instances. Nothing to work with, terminating")
        }
        // Reseting the instance count
        conf.Servers = make([]Instance, 0)
    }

    for i:= 0; i < len(params.Instances);i ++ {
        inst := Instance{Domain: params.Instances[i]}
        conf.Servers = append(conf.Servers, inst)
    }

    return conf
}
