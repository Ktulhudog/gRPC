package sender

import (
	"context"
	sender "gRPC/proto/gogenproto"
)

type GRPCServer struct {
}

func (s *GRPCServer) Send(ctx context.Context, req *sender.SendRequest) (*sender.SendResponse, error) {
	return &sender.SendResponse{Body: req.GetBody()}, nil //dobavit pomenyat req.body
}
