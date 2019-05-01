package config

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// LoadConfiguration - Load the configuration for the service.
func LoadConfiguration() *Service {
	var conf Service
	cliParams := getCliParams()
	// Reading the config file
	configFile := "./config.toml"
	if len(*(cliParams.ConfigFile)) > 0 {
		if _, err := os.Stat(*(cliParams.ConfigFile)); err == nil {
			// There is a file with that name, so use it as a config file
			// It is safe to reuse err as it is nil
			if configFile, err = expand(*(cliParams.ConfigFile)); err != nil {
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
	conf = updateConfig(conf, cliParams)

	return &conf
}

// Updating config from CLI, if needed
func updateConfig(conf Service, params CliParams) Service {
	// Checking the flags

	// -- Key Count --
	if *params.UnsealKeyCount > 0 {
		conf.Vault.UnsealKeyCount = *params.UnsealKeyCount
	}
	// -- Check interval --
	if *params.Interval > 0 {
		conf.Vault.CheckInterval = *params.Interval
	}

	// -- Protocol --
	// converting protocol to lowercase
	prtcl := strings.ToLower(*params.Protocol)
	// and checking if its supported
	if prtcl == "http" || prtcl == "https" {
		conf.Vault.Protocol = *params.Protocol
	} else if len(*params.Protocol) > 0 {
		log.Printf("Unsupported protocol %s given on protocol flag. Using the one defined in the configuration file instead\n", *params.Protocol)
	}

	// -- Instances --
	// Checking that if reset instance is givem, at least one instance is also given
	if *params.ResetInstances {
		if len(params.Instances) == 0 {
			log.Fatal("An instance reset was passed with no new instances. Nothing to work with, terminating")
		}
		// Reseting the instance count
		conf.Servers = make([]Instance, 0)
	}

	for i := 0; i < len(params.Instances); i++ {
		inst := Instance{Domain: params.Instances[i]}
		conf.Servers = append(conf.Servers, inst)
	}

	// -- Pathing --
	if len(*params.StatusPath) > 0 {
		conf.Vault.StatusPath = *params.StatusPath // @TODO check its url legal
	}
	if len(*params.UnsealPath) > 0 {
		conf.Vault.UnsealPath = *params.UnsealPath // @TODO check its url legal
	}

	// -- Keys --
	if len(params.Keys) > 0 {
		// Checking for double entry
		if len(*params.KeyFile) > 0 {
			log.Fatal("Cannot take both a key file and key parameters. Choose one or the other.\nOr better yet. Non of them\n")
		}
		conf.Keys = params.Keys[:]
	}
	if len(*params.KeyFile) > 0 {
		var err error
		// Read the keys of the file. Each line represents a key
		filePath, _ := expand(*params.KeyFile) // Supporting ~ in path
		conf.Keys, err = loadKeyFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load keys from file %s\n", *params.KeyFile)
		}
	}
	return conf
}

func loadKeyFile(filePath string) ([]string, error) {
	// Open the file and put all the keys in a slice
	// Each line in the file is expected to be a key
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s: %s", filePath, err)
		return nil, err
	}
	defer file.Close()

	keys := make([]string, 0)

	// Reading the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Adding key after whitespace cleaning
		keys = append(keys, strings.TrimSpace(scanner.Text()))
	}

	// Checking no errors were raized during reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return keys, err
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
