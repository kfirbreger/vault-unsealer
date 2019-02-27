package unsealer

import (
    "fmt"
    "http"
    "time"
)


type unsealparams struct {
    Key *string
    reset bool
    migrate bool
}


func checkStatus(url) (bool, error) {
    // Checks the vault status
    var unsealed bool
    resp, err := http.Head(url)
    if err != nil {
        fmt.Println(err)
        return (unsealed, err)
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        unsealed = true
    } else {
        // Status not in the 200, something is wrong
        unsealed = false
        err = error('Status code returned: ' + resp.StatusCode)
    }
    return unsealed, err
}

func Unsealer(vault *Vault, instance *Instance, status chan int, keys *[]string) {
    // Using status to send status update to main process
    // Creating the urls to use
    statusUrl := *vault.protocl + "://" + *instance.Domain + *vault.StatusPath
    unsealUrl := *vault.protocol + "://" + *instance.Domain + *vault.UnsealPath

    for {
        status, err := checkStatus(statusUrl)
        if err != nil {
            fmt.Println(err)
        } else {
            if !unsealed {
                // This instance is sealed and needs to be unsealed
                for i := 0;i < len(keys); i++ {
                    

