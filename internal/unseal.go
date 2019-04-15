package internal

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/awnumar/memguard"
)

const UNSEALCALLERROR = -1
const UNSEALHTTPTIMEOUT = 2000

type Unsealparams struct {
	Keys    []*memguard.LockedBuffer
	Reset   bool
	Migrate bool
}

type Unsealer struct {
	ID          int
	UnsealQueue chan UnsealRequest
	ManageChan  chan int
	LogChan     chan<- string
	params      *Unsealparams
}

func NewUnsealer(id int, unsealQueue chan UnsealRequest, logChan chan string, up *Unsealparams) *Unsealer {
	unsealer := Unsealer{
		ID:          id,
		UnsealQueue: unsealQueue,
		LogChan:     logChan,
		params:      up,
	}

	return &unsealer
}

func ExecUnsealOverHttp(id int, key *memguard.LockedBuffer, url string, reset bool, migrate bool) (status int, err error) {
	// Perform an unseal request over http(s)
	// Again key is passed as pointer to prevent leaking to gc
	// Creating a buffer with the key. This is unfortunaltely unavoidable
	// TODO add reset and migrate options to the call
	log.Printf("Unsealer %d - Creating unseal request %s", id, key)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(append([]byte(`{"key":"`), append((*key).Buffer(), []byte(`"}`)...)...)))
	// Sending the request
	timeout := time.Duration(UNSEALHTTPTIMEOUT * time.Millisecond)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	// Making sure body is closed
	if err != nil {
		return UNSEALCALLERROR, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, err
}

func (u *Unsealer) Start() {
	go func() {
		for {
			select {
			case unsealRequest := <-u.UnsealQueue:
				// Performing the unsealing request
				if unsealRequest.KeyNumber >= len(u.params.Keys) { // Sanity check
					// Making sure there is a key available
					log.Printf("Key %d is out of range\n", unsealRequest.KeyNumber)
				}
				log.Println("Unseal request recieved", u.params.Keys)
				status, err := ExecUnsealOverHttp(u.ID, u.params.Keys[unsealRequest.KeyNumber], unsealRequest.Url, u.params.Reset, u.params.Migrate)
				if err != nil {
					log.Println("Error sending unseal call")
				}
				log.Printf("Unseal returned status code %d\n", status)

			case cmd := <-u.ManageChan:
				switch cmd {
				case STOP:
					log.Printf("Stopping unsealer %d", u.ID)
					return
				default:
					log.Printf("Unsealer %d got unknown command %d\n", u.ID, cmd)
				}
            default:
                continue
			}
		}
	}()
}

func (u *Unsealer) Stop() {
	u.ManageChan <- STOP
}
