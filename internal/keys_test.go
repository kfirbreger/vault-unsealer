package internal

import (
	"testing"
)

/*
func TestCreateCorrectKeyCount(t *testing.T) {
    keyCount := 2

    keys := GetUnsealKeys(keyCount, nil)

    if len(keys) != keyCount {
        t.Errorf("Expected %d keys, but instead got %d", keyCount, len(keys))
    }
}
*/
func TestCorrctlyUsingPreloadedKeys(t *testing.T) {
	keyCount := 4
	configKeys := []string{"1", "2", "3", "4"}

	keys := GetUnsealKeys(keyCount, configKeys)

	// Making sure the key value match
	if configKeys[0] != string(keys[0].Buffer()) {
		t.Errorf("Key in memory does not match key given")
	}
}
