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
	}
	reader := bufio.NewReader(os.Stdin)
	var singleKey = make([]byte, 0)

	for i := 0; i < keyCount; i++ {
		if readKeys {
			fmt.Printf("Unsealing key %d: ", i)
			singleKey, _ = reader.ReadBytes('\n')
		} else {
			singleKey = []byte(configKeys[i])
		}
		// convert CRLF to LF
		singleKey = bytes.TrimSpace(singleKey)
		membuf, err := memguard.NewImmutableFromBytes(singleKey)
		if err != nil {
			log.Fatal("Eror creating memory safe storage")
		}
		keys = append(keys, membuf)
	}
	return keys
}
