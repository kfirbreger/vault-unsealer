package internal

import (
	"bufio"
	"bytes"
	"log"
	"os"
	
    "github.com/awnumar/memguard"
)

func GetUnsealKeys(keyCount int) []*memguard.LockedBuffer {
	// Save the unsealing keys in a slice
	// Need to move it to memguard so its safe in memory
	keys := make([]*memguard.LockedBuffer, 0, keyCount)
	reader := bufio.NewReader(os.Stdin)
	for i := 1; i < keyCount+1; i++ {
		log.Printf("Unsealing key %d: ", i)
		text, _ := reader.ReadBytes('\n')
		// convert CRLF to LF
		text = bytes.TrimSpace(text)
		membuf, err := memguard.NewImmutableFromBytes(text)
		if err != nil {
			log.Println("Eror creating memory safe storage")
		}
		keys = append(keys, membuf)
	}
	return keys
}
