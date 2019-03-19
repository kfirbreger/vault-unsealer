package unsealer

import (
    "encoding/json"
    "fmt"
    "http"
    "runtime"
    "time"
) 

const UNSEALCALLERROR -1


type unsealparams struct {
    Keys *[]LockedBuffer `json:"key"`
    reset bool
    migrate bool
}

type Unsealer struct {
    ID int
    UnsealQueue chan UnsealRequest
    ManageChan chan int
    LogChan <- chan string
    params *unsealparams
}


func NewUnsealer(id int, unsealQueue chan UnsealRequest, logChan chan string, *up unsealparams) Unsealer {
    unsealer := Unsealer {
        ID: id
        UnsealQueue : unsealQueue
        LogChan logchan
        params: up
    }

    return unsealer
}

func (u *Unsealer) Start() {
    go func {
        // Wait until there is an unseal request
        for unsealRequest := range u.UnsealQueue {
            // Performing the unsealing request
            if unsealRequest.KeyNumber >= len(u.params.Keys) {. // Sanity check
                // Making sure there is a key available
                fmt.Printf("Key %d is out of range", unsealRequest.KeyNumber)
            }
            status, err := ExecUnsealOverHttp(u.params.Keys[unsealRequest.KeyNumber], unsealRequest.url, u.params.reset, u.params, migrate)
        }
    }
}


func ExecUnsealOverHttp(key *LockedBuffer, url string, reset bool, migrate bool) (status int, err error) {
    // Perform an unseal request over http(s)
    // Again key is passed as pointer to prevent leaking to gc
    get the key -> key.Buffer()
    // Creating a buffer with the key. This is unfortunaltely unavoidable
    jsonBytesPayload := json() // TODO how does this work again?
    
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytesPayload))
    // clearing buffer. which should speed up memory cleaning
    jsonBytesPayload = nil
    
    // Sending the request
    client := &http.Client()
    resp, err := client.Do(req)
    // Making sure body is closed
    defer resp.Body.close()
    if err != nil {
        return UNSEALCALLERROR, err
    }
    defer runtime.GC()  // Manually triggering GC
    return resp.StatusCode, err
}

func checkStatus(url) (bool, error) {
    // Checks the vault status
    var unsealed bool
    resp, err := http.Head(url)
    if err != nil {
        fmt.Println(err)
        return (unsealed, err)
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        unsealed = true
    } else {
        // Status not in the 200, something is wrong
        unsealed = false
        err = error('Status code returned: ' + resp.StatusCode)
    }
    return unsealed, err
}


func instanceMonitor(vault *Vault, instance *Instance, status chan int, keys *[]string) {
    // Using status to send status update to main process
    // Creating the urls to use
    statusUrl := *vault.protocl + "://" + *instance.Domain + *vault.StatusPath
    unsealUrl := *vault.protocol + "://" + *instance.Domain + *vault.UnsealPath

    for {
        status, err := checkStatus(statusUrl)
        if err != nil {
            fmt.Println(err)
        } else if !unsealed {
            // This instance is sealed and needs to be unsealed
            for i := 0;i < len(keys); i++ {
                payload := unsealparams {
                    Key: &(*keys[i])
                    reset: false
                    migrate: false
                }
                go unsealCall(unsealUrl, payload)
            }
        }
        time.Sleep(*vault.CheckInterval * time.Millisecond)            
    }
}

