package main

import (
	"encoding/json"
	"fmt"
)

type MyObject struct {
	Number int
	Word   string
}

func main() {
	object := MyObject{5, "Packt"}
	oJson, _ := json.Marshal(object)
	fmt.Printf("%s\n", oJson)
}
