package config

import (
    "fmt"
    "log"
    "os"
    "os/user"
    "path/filepath"
    "strings"

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
