package internal

import (
	"log"
	"net/http"
	"time"
)

// Constanses for the manage channel
const STOP = 1
const CONTINUE = 2
const PAUZE = 3
const STATUS = 5

// HTTP timeout
const STATUSHTTPTIMEOUT = 2000

type Checker struct {
	ID               int
	StatusCheckQueue chan StatusCheckRequest
	UnsealQueue      chan<- string
	ManageChan       chan int
	LogChan          chan<- string
	CallsMade        int
	CallsSuccessful  int
	UnsealRequests   int
}

func NewChecker(id int, statusCheckQueue chan StatusCheckRequest, unsealQueue chan string, logChan chan string) *Checker {
	// Creating a new worker
	checker := Checker{
		ID:               id,
		StatusCheckQueue: statusCheckQueue,
		UnsealQueue:      unsealQueue,
		ManageChan:       make(chan int),
		LogChan:          logChan,
		CallsMade:        0,
		CallsSuccessful:  0,
		UnsealRequests:   0,
	}

	return &checker
}

// this is set as a differenct function so that
// the actual check can be done differently (like listening to kubernetes api)
// without significant change code
func ExecCheckOverHttp(id int, url string) (int, error) {
	// Makeing a call, returning the status code, or error code
	timeout := time.Duration(STATUSHTTPTIMEOUT * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	// Debuging info
	if err != nil {
		log.Printf("Checker %d - Error calling to Vault.\n", id)
		log.Println(err)
        log.Println("Response is:", resp)
		return -1, err
	}
	log.Printf("Checker %d - StatusCode: %d\n", id, resp.StatusCode)
	return resp.StatusCode, err
}

// @TODO replace all log with log channel publishing
func (c *Checker) Start() {
	go func() {
		for {
			select {
			case work := <-c.StatusCheckQueue:
				// Recieve work request
				log.Printf("worker %d: Received check request for url %s\n", c.ID, work.Url)

				c.CallsMade++
				StatusCode, err := ExecCheckOverHttp(c.ID, work.Url)
				// Checking vault status
				if StatusCode == 503 { // TODO unseal conditions
					c.UnsealRequests++
					c.UnsealQueue <- work.Domain
				} else if StatusCode > 199 && StatusCode < 300 && err == nil {
					c.CallsSuccessful++
				} else {
                    log.Printf("Checker %d - Unknown status code %d\n", c.ID, StatusCode)
                }

			case cmd := <-c.ManageChan:
				log.Printf("Command recieved: %d", cmd)
				switch cmd {
				case STOP:
					log.Printf("Worker %d quitting", c.ID)
					return

				case STATUS:
					log.Printf("Statistics:\n=======================\nCalls made: %d\nCalls succesfull: %d\nUnseal initiated: %d\n", c.CallsMade, c.CallsSuccessful, c.UnsealRequests)

				default:
					log.Printf("Command %d not (yet) supported", cmd)
				}
            default:
			    continue
            }
		}
	}()
}

// Adding worker stop function
func (c *Checker) Stop() {
	c.ManageChan <- STOP
}
