package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/jhsc/say/api"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	port := flag.Int("p", 8080, "listening port")
	flag.Parse()

	fmt.Printf("listening on port %d", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		fmt.Errorf("error listening on port %d: %v", *port, err)
	}

	svr := grpc.NewServer()
	pb.RegisterTextToSpeechServer(svr, server{})

	err = svr.Serve(listener)
	if err != nil {
		fmt.Errorf("could not serve: %v", err)
	}
}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	return nil, nil
}
