package internal

import (
    "fmt"
    "testing"
    "time"
)


const ALLOWEDTIMEDIFFERENCETHRESHOLD = 10


func TestUrlCreatedCorrectly(t *testing.T) {
    sc := make(chan StatusCheckRequest, 1)
    domain := "localhost:1234"
    protocol := "https"
    stsPth := "1/2/3"
    interval := 1000
   
    expectedUrl := protocol + "://" + domain + "/" + stsPth
    go GenerateChecks(sc, domain, protocol, stsPth, interval)
    // Waiting for a message
    msg := <- sc
    
    if msg.Url != expectedUrl {
        t.Errorf("Expected url to be \"https://localhost:1234/1/2/3\" but got %s instead\n", msg.Url)
    }
}

func TestStatusCheckInterval(t *testing.T) {
    sc := make(chan StatusCheckRequest, 1)
    domain := "localhost:1234"
    protocol := "https"
    stsPth := "1/2/3"
    interval := 370
    
    go GenerateChecks(sc, domain, protocol, stsPth, interval)
    // Waiting for a message
    _ = <- sc
    startTime := timestamp()
    
    _ = <- sc
    endTime := timestamp()
    delta := endTime - startTime - int64(interval)
    if delta < -ALLOWEDTIMEDIFFERENCETHRESHOLD || ALLOWEDTIMEDIFFERENCETHRESHOLD < delta {
        t.Errorf("Time interval is used correctly. Expeted ~%d but instead got %d", interval, delta)
    }
}

func TestCreateUnsealRequestPerKey(t *testing.T) {
    un := make(chan string, 10)
    ur := make(chan UnsealRequest, 10)
    protocol := "https"
    unsealPath := "sys/path/to/unseal"
    unsealKeyCount := 5

    domain := "domain.local"

    go GenerateUnseal(un, ur, protocol, unsealPath, unsealKeyCount)
    
    expectedUrl := protocol + "://" + domain + "/" + unsealPath

    // Sending an unseal request
    un <- domain

    requestCounter := 0
    for req := range ur {
        fmt.Println(req)
        if req.Url != expectedUrl {
            t.Errorf("Expecteing an unseal request for %s, but instead got one for %s", req.Url, expectedUrl)
        }
        if req.KeyNumber != requestCounter {
            t.Errorf("Key number in the request should be %d, but instead is %d", requestCounter, req.KeyNumber)
        }

        requestCounter++
    }
    if requestCounter != unsealKeyCount {
        t.Errorf("There should have been %d calls created but instead %d was created", unsealKeyCount, requestCounter)
    }
}

func timestamp() int64 {
    // Generate an int64 timestamp
    return time.Now().UnixNano() / int64(time.Millisecond)
}
