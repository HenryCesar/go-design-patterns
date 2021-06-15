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
	var dangerousObject map[string]interface{}
	err := json.Unmarshal(jsonBytes, &dangerousObject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, ", dangerousObject["number"])
	fmt.Printf("Word is %s\n", dangerousObject["string"])
	fmt.Printf("Error reference is %v\n",
		dangerousObject["nothing"])
}
