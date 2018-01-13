package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/jhsc/say/api"
	"google.golang.org/grpc"
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

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *server, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	// Text from arguments
	text := &pb.Text{Text: flag.Arg(0)}
	resp, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}

	if err := ioutil.WriteFile(*output, resp.Audio, 0666); err != nil {
		log.Fatalf("could not write to %s: %v", *output, err)
	}
}
