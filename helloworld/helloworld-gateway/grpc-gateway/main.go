package main

import (
	"time"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbAPI "github.com/jaeheonji/polyglot-monorepo/protobuf/helloworld/api/v1"
	pbRPC "github.com/jaeheonji/polyglot-monorepo/protobuf/helloworld/rpc"
)

const (
	address = "localhost:50051"
	port = ":50052"
)

type greeterServiceServer struct {
	client pbRPC.RPCServiceClient
}

func (g *greeterServiceServer) SayHello(ctx context.Context, in *pbAPI.HelloRequest) (*pbAPI.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := g.client.SayUnicorn(ctx, &pbRPC.UnicornRequest{Name: in.GetName()})
	if err != nil {
		log.Fatalf("Not Received UnicornReply: %v", err)
	}

	return &pbAPI.HelloReply{Message: r.GetMessage()}, nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()
	client := pbRPC.NewRPCServiceClient(conn)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pbAPI.RegisterGreeterServiceServer(server, &greeterServiceServer{
		client: client,
	})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Printf("gRPC server listening on %s", port)
	server.Serve(lis)
}
