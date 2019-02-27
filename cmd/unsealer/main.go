package main

import (
    "bufio"
    "flag"
    "fmt"
    "http"
    "os"
    "strings"
    "time"

    "github.com/kfirbreger/vault-unsealer/vaultUnsealer"
)


// Setting default unseal key count
const unsealKeyCount = 3
const millisecondCheckDelay = 100  // How long to wait before vault status checks

func getKeyCount() int {
    keyCount := flag.Int("unsealing-keys", unsealKeyCount, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
    flag.Parse()
    return *keyCount
}

func readKeys(keyCount int) []string {
    // Save the unsealing keys in a slice
    // Need to move it to memguard so its safe in memory
    var keys []string
    reader := bufio.NewReader(os.Stdin)
    for i:= 1; i < keyCount + 1; i++ {
        fmt.Printf("Unsealing key %d: ", i)
        text, _ := reader.ReadString('\n')
        // convert CRLF to LF
        text = strings.Replace(text, "\n", "", -1)
        keys = append(keys, text)
    }
    return keys
}

func checkVaultStatus(url string) {
    // Checks what the vault status is
    // Once an http error is returned
    // The vault is sealed and needs to be unsealed
    for {
        _, err := ttp.get(url)
        if err != nil {
            fmt.Println("Vault is sealed. Time to unseal")
        }
        time.Sleep(millisecondCheckDelay * time.Millisecond)
    }
}

func unsealVault(unsealKeys *[]string, vaultUrl *string) {

func main() {
    // var key_count int
    keyCount := getKeyCount()
    fmt.Println("Unsealing key count: ", keyCount)
    keys := readKeys(keyCount)
    fmt.Println(keys)
}
