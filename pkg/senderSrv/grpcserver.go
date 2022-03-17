package serverSender

import (
	"context"
	"gRPC/pkg/gogenproto"
	"github.com/valyala/fasthttp"
	"log"
)

type GRPCServer struct {
	httpClient fasthttp.Client

	sender.UnimplementedSenderServer
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		httpClient: fasthttp.Client{
			TLSConfig: nil,
		},
	}
}

func (s *GRPCServer) Send(ctx context.Context, request *sender.SendRequest) (*sender.SendResponse, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(request.GetUrl())
	req.Header.SetMethod(request.GetMethod())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := s.httpClient.Do(req, resp); err != nil {
		log.Println("error while calling uri")
		return nil, err
	}

	//if statusCode := resp.StatusCode(); statusCode != fasthttp.StatusOK {
	//	log.Printf("status code[%d](%s) is not [200](OK), while calling url", statusCode, fasthttp.StatusMessage(statusCode))
	//	return nil, errors.New("status code is not [200](OK)")
	//}

	str := string(resp.Body())
	return &sender.SendResponse{Body: &str}, nil
}
