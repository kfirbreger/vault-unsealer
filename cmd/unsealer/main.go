package main

import (
	"github.com/kfirbreger/vault-unsealer/internal"
)

// GODEBUG=clobberfree=1

func main() {
	// Load config
	conf := internal.LoadConfiguration()
	// Update with command line arguments
	// @TODO add command line arguments handling

	// Retrieve unseal keys, and put them
	// in a memguard array. Only one copy is
	// allowed to exist to all refrences will
	// be done via pointers to prevent gc from
	// moving the keys around
	keys := internal.GetUnsealKeys()

	// Start the following:
	// 1. Channels
	// 2. Checker workers
	// 3. Unseal workers
	// 4. Status check generators
	checkerQueue := make(chan internal.StatusCheckRequest, 10) // TODO make this service count * something
	unsealQueue := make(chan internal.UnsealRequest, 20)       // TODO same here
	logChan := make(chan string, 10)

	// Creating Cehcker workers
    checkers := make([]*internal.Checker, 0, conf.Workers.StatusCheckCount)
    unsealers := make([]*internal.Unsealer, 0, conf.Workers.UnsealCount)
    // loggers := make([]string, 0, conf.Workers.LoggingCount) // TODO add loggers

	for i := 0; i < conf.Workers.StatusCheckCount; i++ {
		// Creating checkers
		c := internal.NewChecker(i, checkerQueue, unsealQueue, logChan)
		(*c).Start()
		checkers[i] = c
	}

	// Creating unseal params
	up := &internal.Unsealparams{keys, false, false}
	// Creating unsealer workers
	for i := 0; i < conf.Workers.UnsealCount; i++ {
        u := internal.NewUnsealer(i, unsealQueue, logChan, up)
		(*u).Start()
		unsealers[i] = u
	}

	// Creating the Status check generators
	for i := 0; i < len(conf.Servers); i++ {
		go internal.GenerateChecks(checkerQueue, conf.Servers[i].Domain, conf.Vault.Protocol, conf.Vault.StatusPath, conf.Vault.CheckInterval)
	}
	// Just let the program do its work
	for {
	}
}
