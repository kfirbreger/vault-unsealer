package unsealer

import (
    "fmt"
    "time"
)


func GenerateChecks(<-chan StatusChecks, domain string, protocol string, statusPath string, interval int) {
    // Making a map with the params
    url := protocol + '://' + domain + '/' + statusPath
    // Defining the other parameters
    name := "Status check for " + domain
    for {
        // Creating work
        work := CheckRequest{Name: name, Url: url}
        // Adding it to the work queue
        StatusChecks <- work
        // Waiting before next call
        time.Sleep(interval)
    }
}


func GenerateUnseal(<-chan UnsealRequest, domain string, protocol string, unsealPath string, unsealKeyCount) {
    // Making a map with the params
    params := map[string]string{
        "url": protocol + '://' + domain + '/' + statusPath
    }
    // Defining the other parameters
    name := "Status check for " + domain
    for i := 0; i < unsealKeyCount; i++ {
        // Creating work
        work := UnsealRequest{Name: name, Url: url}
        // Adding it to the work queue
        StatusChecks <- work
        // Waiting before next call
        time.Sleep(interval)
    }
}

