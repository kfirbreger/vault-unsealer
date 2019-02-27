package unsealer

import (
    "fmt"
    "http"
)


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
