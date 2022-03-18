package main

import (
	"context"
	sender "gRPC/server/pkg/gogenproto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	requestData := new(sender.SendRequest)
	requestData.Url = "http://172.16.122.62:30845/api"
	requestData.Method = "POST"

	client := sender.NewSenderClient(connection)
	res, err := client.Send(context.Background(), requestData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetBody())
}
