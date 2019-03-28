package internal

import (
    "testing"
    "time"
)


const ALLOWEDTIMEDIFFERENCETHRESHOLD 10


func TestUrlCreatedCorrectly(t *testing.T) {
    sc := make(chan StatusCheckRequest, 10)
    domain := "localhost:1234"
    protocol := "https"
    stsPth := "1/2/3"
    interval := 1000
    
    go GenerateChecks(sc, domain, protocol, stsPth, interval)
    // Waiting for a message
    msg <- sc
    
    if msg.Url != "https://localhost:1234/1/2/3" {
        t.Errorf("Expected url to be \"https://localhost:1234/1/2/3\" but got %s instead\n", msg.Url)
    }
}

func TestStatusCheckInterval(t *testing.T) {
    sc := make(chan StatusCheckRequest, 10)
    domain := "localhost:1234"
    protocol := "https"
    stsPth := "1/2/3"
    interval := 100
    
    go GenerateChecks(sc, domain, protocol, stsPth, interval)
    // Waiting for a message
    msg <- sc
    startTime := timestamp()
    
    msg <- sc
    endTime := timestamp()
    
    if (endTime - startTime - interval) < ALLOWEDTIMEDIFFERENCETHRESHOLD
}

func timestamp() int64 {
    // Generate an int64 timestamp
    return time.Now().UnixNano() / int64(time.Millisecond)
}