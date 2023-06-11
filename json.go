package main

import (
	"encoding/json"
	"fmt"
)

func jsonify(x any) []byte {
	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	return bytes
}

func RunJsonifyDemo() {
	str := "what is this?"
	fmt.Println(string(jsonify(str)))
}
