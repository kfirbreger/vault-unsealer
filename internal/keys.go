package internal

import (
    "bufio"
    "flag"
	"fmt"
    "os"
    "strings"
	"github.com/awnumar/memguard"
)

func GetUnsealKeys() *[]memguard.LockedBuffer {
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
		text, _ := reader.ReadString("\n")
		text = strings.Replace(text, "\n", "", -1)
		keyCount = int(text)
	}
	return readKeys(keyCount)
}

func getKeyCount() int {
	keyCount := flag.Int("unsealing-keys", 0, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
	flag.Parse()
	return *keyCount
}

func readKeys(keyCount int) *[]memguard.LockedBuffer {
	// Save the unsealing keys in a slice
	// Need to move it to memguard so its safe in memory
	var keys [keyCount]*memguard.LockedBuffer
	reader := bufio.NewReader(os.Stdin)
	for i := 1; i < keyCount+1; i++ {
		fmt.Printf("Unsealing key %d: ", i)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		membuf := memguard.NewImmutableFromBytes(text)
		keys = append(keys, &membuf)
	}
	return &keys
}
