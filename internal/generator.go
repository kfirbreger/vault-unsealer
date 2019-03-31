package internal

import (
	"fmt"
	"time"
    "strconv"
)

// After generating unseal calls, how long shall the generator pause
// before looking for other unsealing calls
const UNSEALSLEEP = 100

func GenerateChecks(statusCheck chan<- StatusCheckRequest, quitChan chan bool, domain string, protocol string, statusPath string, interval int) {
	// Making a map with the params
	url := protocol + "://" + domain + "/" + statusPath
	// Defining the other parameters
	name := "Status check for" + domain
	for {
        select {
        case <- quitChan:
            // Stopping
            return
        default:
            // No quit signal, doing work
            fmt.Println("Checking status of", domain)
            // Creating work
            work := StatusCheckRequest{Name: name, Url: url, Domain: domain}
            // Adding it to the work queue
            statusCheck <- work
            // Waiting before next call
        }
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func GenerateUnseal(unsealNeeded <-chan string, quitChan chan bool, unsealRequest chan<- UnsealRequest, protocol string, unsealPath string, unsealKeyCount int) {
    for { 
        select {
        case <- quitChan:
            // Terminating
            return
        case domain := <- unsealNeeded:
            url := protocol + "://" + domain + "/" + unsealPath
            // Defining the other parameters
            for i := 0; i < unsealKeyCount; i++ {
                name := "Unseal request for " + domain + " with key " + strconv.Itoa(i)
                // Creating work
                work := UnsealRequest{Name: name, Url: url, KeyNumber: i}
                // Adding it to the work queue
                unsealRequest <- work
            }
        }
		time.Sleep(time.Duration(UNSEALSLEEP) * time.Millisecond)
	}
}
