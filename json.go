package main

import (
	"encoding/json"
)

func jsonify(x any) []byte {
	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	return bytes
}
