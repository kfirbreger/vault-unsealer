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
    interval := 370
    
    go GenerateChecks(sc, domain, protocol, stsPth, interval)
    // Waiting for a message
    msg <- sc
    startTime := timestamp()
    
    msg <- sc
    endTime := timestamp()
    delta := endTime - startTime - interval
    if delta < -ALLOWEDTIMEDIFFERENCETHRESHOLD || ALLOWEDTIMEDIFFERENCETHRESHOLD < delta {
        t.Errorf("Time interval is used correctly. Expeted ~%d but instead got %d", interval, delta)
    }
}

func timestamp() int64 {
    // Generate an int64 timestamp
    return time.Now().UnixNano() / int64(time.Millisecond)
}
