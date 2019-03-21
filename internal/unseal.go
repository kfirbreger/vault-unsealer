package internal

import (
    "bytes"
	"fmt"
	"net/http"
	"runtime"
	"github.com/awnumar/memguard"
)

const UNSEALCALLERROR = -1

type unsealparams struct {
	Keys    []*memguard.LockedBuffer 
	reset   bool
	migrate bool
}

type Unsealer struct {
	ID          int
	UnsealQueue chan UnsealRequest
	ManageChan  chan int
	LogChan     chan<- string
	params      *unsealparams
}

func NewUnsealer(id int, unsealQueue chan UnsealRequest, logChan chan string, up *unsealparams) *Unsealer {
	unsealer := Unsealer{
		ID:          id,
		UnsealQueue: unsealQueue,
		LogChan:     logChan,
		params:      up,
	}

	return &unsealer
}

func (u *Unsealer) Start() {
	go func() {
		// Wait until there is an unseal request
		for unsealRequest := range u.UnsealQueue {
			// Performing the unsealing request
			if unsealRequest.KeyNumber >= len(u.params.Keys) { // Sanity check
				// Making sure there is a key available
				fmt.Printf("Key %d is out of range", unsealRequest.KeyNumber)
			}
			status, err := ExecUnsealOverHttp(u.params.Keys[unsealRequest.KeyNumber], unsealRequest.Url, u.params.reset, u.params.migrate)
            if err != nil {
                fmt.Println("Error sending unseal call")
            }
            fmt.Printf("Unseal returned status code %d", status)
		}
	}()
}

func ExecUnsealOverHttp(key *memguard.LockedBuffer, url string, reset bool, migrate bool) (status int, err error) {
	// Perform an unseal request over http(s)
	// Again key is passed as pointer to prevent leaking to gc
	
    // Creating a buffer with the key. This is unfortunaltely unavoidable
    // TODO add reset and migrate options to the call
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(append([]byte(`{"key":"`), append((*key).Buffer(), []byte(`"}`)...)...)))

	// Sending the request
	client := &http.Client{}
	resp, err := client.Do(req)
	// Making sure body is closed
	defer resp.Body.Close()
	if err != nil {
		return UNSEALCALLERROR, err
	}
	defer runtime.GC() // Manually triggering GC
	return resp.StatusCode, err
}
