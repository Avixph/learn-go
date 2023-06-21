package main

import (
	"encoding/json"
	"fmt"
)

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
	}

	return bs, nil
}
