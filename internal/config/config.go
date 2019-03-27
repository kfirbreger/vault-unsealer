package config

import (
    "log"
    "os"
)
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
    // Update with the CLI params
    updateConfig(conf, cliParams)

	return &conf
}

