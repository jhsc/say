package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"

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
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}
