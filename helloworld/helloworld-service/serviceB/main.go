package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/jaeheonji/polyglot-monorepo/protobuf/helloworld/rpc"
)

const (
	port = ":50051"
)

type rpcServiceServer struct {}

// SayUnicorn implements helloworld.rpc.RPCService.SayUnicorn
func (r *rpcServiceServer) SayUnicorn(ctx context.Context, in *pb.UnicornRequest) (*pb.UnicornReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.UnicornReply{Message: "🦄 Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRPCServiceServer(server, &rpcServiceServer{})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Printf("gRPC server listening on %s", port)
	server.Serve(lis)
}
