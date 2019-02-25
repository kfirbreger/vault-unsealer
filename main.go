package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    unseal_count := flag.Int("unsealing-keys", 3, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
    flag.Parse()
    fmt.Println("Unsealing key count: ", *unseal_count)

    // Save the unsealing keys in a slice
    var keys []string
    for i:= 1; i < *unseal_count + 1; i++ {
        fmt.Printf("Unsealing key %d: ", i)
        text, _ := reader.ReadString('\n')
        // convert CRLF to LF
        text = strings.Replace(text, "\n", "", -1)
        keys = append(keys, text)
    }
    fmt.Println(keys)
}
