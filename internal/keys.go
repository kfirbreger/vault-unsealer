package unsealer


func getKeyCount() int {
    keyCount := flag.Int("unsealing-keys", nil, "The number of keys that are required to unseal the vault. You will be prompt for them after this")
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
