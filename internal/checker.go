package internal

import (
    "fmt"
    "http"
)

// Constanses for the manage channel
const QUIT = -1
const CONTINUE = 2
const PAUZE = 3
const STATUS = 5

type Checker struct {
    ID int
    StatusCheckQueue chan StatusCheckRequest
    UnsealQueue chan UnsealRequest
    ManageChan chan int
    LogChan <-chan string
    CallsMade int
    CallsSuccesful int
    UnsealRequests int
}

func NewChecker(id int, statusCheckQueue chan StatusCheckRequest, unsealQueue chan UnsealRequest, logChan chan string) *Checker {
    // Creating a new worker
    checker := Checker{
        ID: id,
        StatusCheckQueue: statusCheckQueue,
        ManageChan: make(chan int)
        LogChan: logChan
        CallsMade: 0
        CallsSuccessful: 0
        UnsealRequests: 0
    }

    return &checker
}

// this is set as a differenct function so that
// the actual check can be done differently (like listening to kubernetes api)
// without significant change code
func ExecCheckOverHttp(url) (int, err) {
    // Makeing a call, returning the status code, or error code
    resp, err = http.get(url)
    // Debuging info
    if err != nil {
        fmt.Println('Error calling to Vault. is Vault sealed?')
        fmt.Println(err)
    }
    fmt.Println('StatusCode: %d', resp.StatusCode)
    return resp.StatusCode, err
}

// @TODO replace all fmt with log channel publishing
func (c *Checker) Start() {
    go func() {
        for {
            // Making worker available
            c.WorkerQueue <- c.Work
            
            select {
            case work := <- c.Work:
                // Recieve work request
                fmt.Printf("worker%d: Received check request for url %s", c.ID, work.Url)
                
                c.CallsMade++
                StatusCode, err = ExecCheckOverHttp(work.Url)
                // Checking vault status
                if ====SEAL CONDITIONS==== {
                    c.UnsealRequest++
                    // @TODO generate unseal work
                } else if StatusCode > 199 && StatusCode < 300 && err == nil {
                    c.CallsSuccesful++
                }

            case cmd := <- w.ManageChan:
                fmt.Printf("Command recieved: %d", cmd)
                switch cmd {
                case QUIT:
                    fmt.Printf("Worker %d quitting", w.ID)
                    return

                case STATUS:
                    fmt.Printf("Statistics:\n=======================\nCalls made: %d\nCalls succesfull: %d\nUnseal initiated: %d\n", c.CallsMade, c.CallsSuccesful, ic.UnsealRequests)

                default:
                    fmt.Printf("Command %d not (yet) supported", cmd)
                }
            }
        }
    }
}


// Adding worker stop function
func (c *Checker) Stop() {
    c.ManageChan <- QUIT
}

