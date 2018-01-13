package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	server := flag.String("s", "localhost:8080", "say server address")
	output := flag.String("o", "output.wav", "wav file name")

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("say server address: %s", *server)
	fmt.Printf("file output name: %s", *output)
}
