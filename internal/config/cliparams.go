package config

import (
    "flag"
    "log"
    "strings"
)

// Setting up for multiple insatnces option
type vaultInstances []string

func (v *vaultInstances) String() string {
    // String representation
    return strings.Join((*v)[:], ", ")
}

func (v *vaultInstances) Set(value string) error {
    *v = append(*v, value)
    return nil
}

type CliParams struct {
    ConfigFile     string
	UnsealKeyCount int
    Instances      vaultInstances
    ResetInstances bool
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
    params.ResetInstances = *flag.Bool("reset-instances", false, "Clears out all the instances in the config. This flag can only be used in combination with at least one instance flag")
    // To support multiple values on the instance flag the bindingis slightly different
    flag.Var(&params.Instances, "instance", "Give the domain of an instance")

	flag.Parse()
    log.Println(params)
	return params
}


