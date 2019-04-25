package internal

import (
	"bufio"
	"bytes"
    "fmt"
	"log"
	"os"

	"github.com/awnumar/memguard"
)

func GetUnsealKeys(keyCount int, configKeys []string) []*memguard.LockedBuffer {
	// Save the unsealing keys in a slice
	keys := make([]*memguard.LockedBuffer, 0, keyCount)
	readKeys := true
	// Checking if keys are given through config
	// and if they are, skip reading from user
	if len(configKeys) > 0 {
		readKeys = false
		log.Println("Keys given via CLI.")
	} else {
        fmt.Printf("Unsealing keys needed. You will be asked for %d keys\n", keyCount)
    }

	reader := bufio.NewReader(os.Stdin)
	var singleKey = make([]byte, 0)
	var err error
    
	for i := 0; i < keyCount; i++ {
		if readKeys {
            fmt.Printf("Unseal key #%d: ", (i + 1))
			singleKey, err = reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}
		} else {
			singleKey = []byte(configKeys[i])
		}
		// convert CRLF to LF
		singleKey = bytes.TrimSpace(singleKey)
		membuf, err := memguard.NewImmutableFromBytes(singleKey)
		if err != nil {
			log.Fatal("Eror creating memory safe storage", err)
		}
		keys = append(keys, membuf)
	}
	return keys
}
