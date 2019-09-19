package main

import (
	context "context"
	"fmt"
	"net"
	"os"

	"github.com/yukpiz/go-proto-validators-example/pb"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := grpc.NewServer()

	h := &Handler{}
	pb.RegisterExampleAPIServer(server, h)
	reflection.Register(server)

	p, err := net.Listen("tcp", ":1111")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "===> Starting gRPC Server...")
	if err := server.Serve(p); err != nil {
		panic(err)
	}
}

type Handler struct{}

func (*Handler) Hello(ctx context.Context, req *pb.ExampleMessage) (*pb.Empty, error) {
	fmt.Fprintf(os.Stdout, "Hello!!")
	return nil, nil
}
