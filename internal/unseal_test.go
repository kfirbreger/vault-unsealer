package internal

import (
    "testing"

	"github.com/awnumar/memguard"
)


func TestUnsealingRequest(t *testing.T) {
    id := 0
    uq := make(chan UnsealRequest, 10)
    lc := make(chan string, 10)

    // Generating fake keys
    keys := make([]*memguard.LockedBuffer, 0)
    for i := 0; i < 3; i++ {
        key, _ := memguard.NewMutableRandom(8)
        keys = append(keys, key)
    }

    up := Unsealparams{
        Keys: keys,
        Reset: false,
        Migrate: false,
    }
    // Starting an unsealer
    unsealer := NewUnsealer(id, uq, lc, &up)
    unsealer.Start()

    unsealRequest := UnsealRequest{
        Name: "Testing unseal",
        Url: "http://localhost:5656/unseal",
        KeyNumber: 1,
    }

    uq <- unsealRequest

    unsealer.Stop()
}
