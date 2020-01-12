package main

import (
	context "context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// GRPCServer - server struct
type GRPCServer struct{}

// SayHello - print hello method
func (g *GRPCServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Result: "Hello " + req.Name}, nil
}

func main() {
	server := grpc.NewServer()
	gs := new(GRPCServer)

	RegisterHelloServer(server, gs)
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	log.Println("Listen port 5000")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
