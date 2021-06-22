package main

import (
	"flag"
	"log"
)

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()

	var activeStrategy OutputStrategy

	switch *output {
	case "console":
		activeStrategy = &TextSquare{}
	case "image":
		activeStrategy = &ImageSquare{"/tmp/image.jpg"}
	default:
		activeStrategy = &TextSquare{}
	}

	err := activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
