package internal

import (
    "bytes"
    "bufio"
    "flag"
	"fmt"
    "os"
    "strconv"
    "strings"
	"github.com/awnumar/memguard"
)

func GetUnsealKeys() []*memguard.LockedBuffer {
	/*
	   Retrieve the unsealing keys by first getting the key count
	   and then prompting for the keys one at a time
	   The keys are saved in a LockedBuffer, a struct from
	   the memguard package
	*/

	// Getting key count
	keyCount := getKeyCount()
	// If no key parameter is given, requesting key count
	// This might change to fall back to config first
	if keyCount == 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Unsealing key count: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		keyCount, _ = strconv.Atoi(text)
	}
	return readKeys(keyCount)
}

func getKeyCount() int {
	keyCount := flag.Int("unsealing-keys", 0, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
	flag.Parse()
	return *keyCount
}

func readKeys(keyCount int) []*memguard.LockedBuffer {
	// Save the unsealing keys in a slice
	// Need to move it to memguard so its safe in memory
    keys := make([]*memguard.LockedBuffer, keyCount)
	reader := bufio.NewReader(os.Stdin)
	for i := 1; i < keyCount+1; i++ {
		fmt.Printf("Unsealing key %d: ", i)
		text, _ := reader.ReadBytes('\n')
		// convert CRLF to LF
		text = bytes.TrimSpace(text)
		membuf, err := memguard.NewImmutableFromBytes(text)
        if err != nil {
            fmt.Println("Eror creating memory safe storage")
        }
		keys = append(keys, membuf)
	}
	return keys
}
