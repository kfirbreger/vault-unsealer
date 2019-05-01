package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// A type for the instance url
type Instance struct {
	Domain string
}

// The Vault configuration
// How to call, where to call, and how often to call
type VaultConf struct {
	Protocol       string
	StatusPath     string `toml:"status_path"`
	UnsealPath     string `toml:"unseal_path"`
	UnsealKeyCount int    `toml:"unseal_key_count"`
	CheckInterval  int    `toml:"seal_check_interval"`
}

// The workers configuration
type WorkersConf struct {
	StatusCheckCount int `toml:"status_check"`
	UnsealCount      int `toml:"unseal"`
	LoggingCount     int `toml:"logging"`
}

// The overarching configuration type
type Service struct {
	Vault   VaultConf
	Workers WorkersConf
	Servers []Instance `toml:"server"`
	Keys    []string
}

// Load the the configuration file
func Load(filepath string, s *Service) {
	if _, err := toml.DecodeFile(filepath, s); err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration: ", *s)
	if len((*s).Keys) > 0 {
		log.Fatal("No keys are allowed in the configuration file")
	}
	// @TODO add verification
}
