package internal

import (
	"log"
	"net/http"
)

// Constanses for the manage channel
const STOP = 1
const CONTINUE = 2
const PAUZE = 3
const STATUS = 5

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
func ExecCheckOverHttp(url string) (int, error) {
	// Makeing a call, returning the status code, or error code
    timeout := time.Duration(200 * time.Millisecond)
    client := http.Client{
        Timeout: timeout,
    }
    resp, err := client.Get(url)
	// Debuging info
	if err != nil {
		log.Println("Error calling to Vault. is Vault sealed?")
		log.Println(err)
		return -1, err
	}
	log.Println("StatusCode: ", resp.StatusCode)
	return resp.StatusCode, err
}

// @TODO replace all fmt with log channel publishing
func (c *Checker) Start() {
	go func() {
		for {
			select {
			case work := <-c.StatusCheckQueue:
				// Recieve work request
				fmt.Printf("worker %d: Received check request for url %s\n", c.ID, work.Url)

				c.CallsMade++
				StatusCode, err := ExecCheckOverHttp(work.Url)
				// Checking vault status
				if StatusCode == 503 { // TODO unseal conditions
					c.UnsealRequests++
					c.UnsealQueue <- work.Domain
				} else if StatusCode > 199 && StatusCode < 300 && err == nil {
					c.CallsSuccessful++
				}

			case cmd := <-c.ManageChan:
				fmt.Printf("Command recieved: %d", cmd)
				switch cmd {
				case STOP:
					fmt.Printf("Worker %d quitting", c.ID)
					return

				case STATUS:
					fmt.Printf("Statistics:\n=======================\nCalls made: %d\nCalls succesfull: %d\nUnseal initiated: %d\n", c.CallsMade, c.CallsSuccessful, c.UnsealRequests)

				default:
					fmt.Printf("Command %d not (yet) supported", cmd)
				}
			}
		}
	}()
}

// Adding worker stop function
func (c *Checker) Stop() {
	c.ManageChan <- STOP
}
