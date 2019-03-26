package internal

import (
    "log"
    "os/exec"
    "testing"
    "time"
)

func setupDevVault() {
    /**
    Start vault in dev mode to test that the checker is working correctly
    */
    cmd := exec.Command("../testing/vault-dev.sh")
    err := cmd.Start()
    if err != nil {
        log.Printf("Unable to start vault in dev mode. Testing will fail\n")
        log.Fatal(err)
    }
    // Waiting for the docker run command to return, and give vault time to init
    err = cmd.Wait()
    time.Sleep(time.Duration(500) * time.Millisecond)
    if err != nil {
        log.Fatal(err)
    }
}

func teardownDevVault() {
    /**
    Stopping the server
    */
    cmd := exec.Command("docker", "stop", "vault-dev") // @TODO exec pathing?
    err := cmd.Start()
    if err != nil {
        log.Println("Failed to stop vault-dev container")
        log.Println(err)
    }
}

func TestChecker(t *testing.T) {
    // Starting by seting up the server
    setupDevVault()
    // Closing down on at the end
    defer teardownDevVault()

    // Creating the needed channels
	checkerQueue := make(chan StatusCheckRequest, 10) 
    unsealNeededQueue := make(chan string, 10)
    logChan := make(chan string, 10)

    // Creating a worker

    c := NewChecker(0, checkerQueue, unsealNeededQueue, logChan)
    (*c).Start()
    
    // Creating a check request
    url := "http://127.0.0.1:8200/v1/sys/health?standbyok=true"
    name := "Testing checker"
    domain := "127.0.0.1"
    work := StatusCheckRequest{Name: name, Url: url, Domain: domain}
	
    // Adding it to the work queue
    checkerQueue <- work
    // Checking that a check was made
    time.Sleep(time.Duration(100) * time.Millisecond)
    if (*c).CallsMade != 1 {
        t.Fatalf("Expected a call to be made but %d calls were made", (c).CallsMade)
    }
    if (*c).CallsSuccessful != 1 {
        t.Fatalf("Excpected one successful call but got %d", (*c).CallsSuccessful)
    }
}
