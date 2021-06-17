package main

import (
	"encoding/json"
	"fmt"
)

type MyObject struct {
	Number int    `json:"number"`
	Word   string `json:"string"`
}

func main() {
	jsonBytes := []byte(`{"number":5, "string":"Packt"}`)
	var object MyObject
	err := json.Unmarshal(jsonBytes, &object)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, Word is %s\n", object.Number, object.Word)
}
