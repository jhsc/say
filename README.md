# Text speech server with gRPC.

Requirements:
- Install [protobuf](https://github.com/google/protobuf/releases)
- Docker

References:
- flite [docs](http://www.speech.cs.cmu.edu/flite/)
- gRPC [docs](https://grpc.io)

## Run server
```
$ docker run --rm -p 8080:8080 jhsc/say-server
```

## Run Client
```
$ go run say/main.go -s localhost:8080 -o output.wav
```

### To do
- Docker-compose
