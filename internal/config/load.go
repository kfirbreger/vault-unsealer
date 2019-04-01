package config

import (
	"log"

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
	Keys    []string
}

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

