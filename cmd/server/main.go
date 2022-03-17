package main

import (
	senderPb "gRPC/pkg/gogenproto"
	serverSender "gRPC/pkg/senderSrv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	senderPb.RegisterSenderServer(grpcServer, serverSender.NewGRPCServer())

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
