package unsealer

import (
    "http"
)

// Constanses for the manage channel
const QUIT = -1
const CONTINUE = 2
const PAUZE = 3
const STATUS = 5

type Worker struct {
    ID int
    Work chan WorkRequest
    WorkerQueue chan chan WorkRequest
    ManageChan chan int
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
    // Creating a new worker
    worker := Worker{
        ID: id,
        Work: make(chan WorkRequest),
        WorkerQueue: workerQueue,
        ManageChan: make(chan int)
    }

    return worker
}

func (w *Worker) call(url) (int, err) {
    // Makeing a call, returning the status code, or error code
    resp, err = http.get(url)
    // Debuging info
    if err != nil {
        fmt.Println('Vault is sealed. Needs unsealing')
        fmt.Println(err)
    }
    fmt.Println('StatusCode: %d', resp.StatusCode)
    return resp.StatusCode, err
}

func (w *Worker) Start() {
    go func() {
        // Some statistics
        callsMade := 0
        callsSuccesful := 0
        unsealedInitiated := 0
        for {
            // Making worker available
            w.WorkerQueue <- w.Work
            
            select {
            case work := <- w.Work:
            // Recieve work request
            fmt.Printf("worker%d: Received work request", w.ID)
            callsMade++

            StatusCode, err = w.call(work.Url)
            // Checking vault status
            if ====SEAL CONDITIONS==== {
                unsealedInitiated++
                @TODO generate unseal work
            } else if StatusCode > 199 && StatusCode < 300 && err == nil {
                callsSuccesful++
            }

            case cmd := <- w.ManageChan:
            fmt.Printf("Command recieved: %d", cmd)
            switch {
            case cmd == QUIT:
                fmt.Printf("Worker %d quitting", w.ID)
                return

            case cmd == STATUS:
                fmt.Printf("Statistics:\n=======================\nCalls made: %d\nCalls succesfull: %d\nUnseal initiated: %d\n", callsMade, callsSuccesful, unsealedInitiated)

            default:
                fmt.Printf("Command %d not (yet) supported", cmd)
            }
            }
        }
    }
}


// Adding worker stop function
func (w *Worker) Stop() {
    w.ManageChan <- QUIT
}

