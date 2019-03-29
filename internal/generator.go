package internal

import (
	"fmt"
	"time"
)

// After generating unseal calls, how long shall the generator pause
// before looking for other unsealing calls
const UNSEALSLEEP = 100

func GenerateChecks(statusCheck chan<- StatusCheckRequest, domain string, protocol string, statusPath string, interval int) {
	// Making a map with the params
	url := protocol + "://" + domain + "/" + statusPath
	// Defining the other parameters
	name := "Status check for" + domain
	for {
		fmt.Println("Generating unseal request for ", domain)
		// Creating work
		work := StatusCheckRequest{Name: name, Url: url, Domain: domain}
		// Adding it to the work queue
		statusCheck <- work
		// Waiting before next call
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func GenerateUnseal(unsealNeeded <-chan string, unsealRequest chan<- UnsealRequest, protocol string, unsealPath string, unsealKeyCount int) {
	for domain := range unsealNeeded {
		url := protocol + "://" + domain + "/" + unsealPath
		// Defining the other parameters
		name := "Unseal request for " + domain
		for i := 0; i < unsealKeyCount; i++ {
			// Creating work
			work := UnsealRequest{Name: name, Url: url, KeyNumber: i}
			// Adding it to the work queue
			unsealRequest <- work
		}
		time.Sleep(time.Duration(UNSEALSLEEP) * time.Millisecond)
	}
}
