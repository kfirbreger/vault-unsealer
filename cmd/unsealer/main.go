package main

import (
	"log"
    "time"

	"github.com/kfirbreger/vault-unsealer/internal"
	"github.com/kfirbreger/vault-unsealer/internal/config"
)

// GODEBUG=clobberfree=1

func main() {
	// Load config
	conf := config.LoadConfiguration()
	// Update with command line arguments
	// @TODO add command line arguments handling

	// Retrieve unseal keys, and put them
	// in a memguard array. Only one copy is
    // allowed to exist to all references will
	// be done via pointers to prevent gc from
	// moving the keys around
	keys := internal.GetUnsealKeys(conf.Vault.UnsealKeyCount, conf.Keys)
	// Start the following:
	// 1. Channels
	// 2. Checker workers
	// 3. Unseal workers
	// 4. Status check generators
	checkerQueue := make(chan internal.StatusCheckRequest, 100) // TODO make this service count * something
	unsealQueue := make(chan internal.UnsealRequest, 50)        // TODO same here
	unsealNeededQueue := make(chan string, 50)
	logChan := make(chan string, 10)
	quitChan := make(chan bool, 10)

	// Creating Cehcker workers
	checkers := make([]*internal.Checker, 0, conf.Workers.StatusCheckCount)
	unsealers := make([]*internal.Unsealer, 0, conf.Workers.UnsealCount)
	// loggers := make([]string, 0, conf.Workers.LoggingCount) // TODO add loggers

	for i := 0; i < conf.Workers.StatusCheckCount; i++ {
		// Creating checkers
		c := internal.NewChecker(i, checkerQueue, unsealNeededQueue, logChan)
		(*c).Start()
		log.Printf("Created %d checker\n", i)
		checkers = append(checkers, c)
	}

	// Creating unseal params
	up := &internal.Unsealparams{keys, false, false}
	// Creating unsealer workers
	for i := 0; i < conf.Workers.UnsealCount; i++ {
		u := internal.NewUnsealer(i, unsealQueue, logChan, up)
		(*u).Start()
		log.Printf("Created %d unsealer\n", i)
		unsealers = append(unsealers, u)
	}

	// Creating the Status check generators
	for i := 0; i < len((*conf).Servers); i++ {
		go internal.GenerateChecks(checkerQueue, quitChan, conf.Servers[i].Domain, conf.Vault.Protocol, conf.Vault.StatusPath, conf.Vault.CheckInterval)
		log.Printf("Created generator for %s\n", conf.Servers[i].Domain)
	}
	go internal.GenerateUnseal(unsealNeededQueue, quitChan, unsealQueue, conf.Vault.Protocol, conf.Vault.UnsealPath, len(keys))
	log.Println("Monitoring started")
	// Just let the program do its work
	for {
        time.Sleep(1 * time.Second)
	}
}
