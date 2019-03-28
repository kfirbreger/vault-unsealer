package internal

import (
    "testing"
    "time"
)

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
