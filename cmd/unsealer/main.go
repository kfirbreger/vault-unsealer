package main

import (
    "os"
    "sync"
    "github.com/kfirbreger/vault-unsealer/unsealer"
)

// GODEBUG=clobberfree=1

func main() {
    // Load config
    conf := unsealer.LoadConfiguration()
    // Update with command line arguments
    // @TODO add command line arguments handling

    // Retrieve unseal keys, and put them
    // in a memguard array. Only one copy is
    // allowed to exist to all refrences will
    // be done via pointers to prevent gc from
    // moving the keys around
    keys := unsealer.GetUnsealKeys()

    // Start the following:
    // 1. Channels
    // 2. Checker workers
    // 3. Unseal workers
    // 4. Status check generators
    checkerQueue := make(chan unsealer.StatusCheckRequest, 10) // TODO make this service count * something
    unsealQueue := make(chan unsealer.UnsealRequest, 20)  // TODO same here
    logChan := make(chan string, 10)

    // Creating Cehcker workers
    var checkers = [conf.Workers.StatusCheckCount]Checker
    var unsealers = [conf.Workers.UnsealCount]Unsealer
    var loggers = [conf.Workers.LoggingCount]string  // TODO add loggers

    for i := 0; i < len(checkers): i++ {
        // Creating checkers
        c := NewChecker(i, checkerQueue, unsealQueue, logChan)
        *c.Start()
        checkers[i] = c
    }

    // Creating unseal params
    up := *unsealer.unsealparams{*keys, false, false}
    // Creating unsealer workers
    for i := 0; i < len(unsealQueue), i++ {
        u = NewUnsealer(i, unsealQueue, logChan, up)
        *u.Start()
        unsealers[i] = u
    }

    // Creating the Status check generators
    for i := 0; i < len(conf.Servers); i++ {
        go unsealer.GenerateChecks(checkerQueue, conf.Servers[i].Domain, conf.Vault.Protocol, conf.Vault.StatusPath, conf.Vault.CheckInterval)
    }
    // Just let the program do its work
    for { }
}

