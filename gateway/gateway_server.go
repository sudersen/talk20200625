package gateway

import (
	"context"
	"fmt"
	"time"
)

// START OMIT
func NewGateway() *gatewayServer { // HLxxx
	return &gatewayServer{}
}

type gatewayServer struct{} // HLxxx

func (s *gatewayServer) Read(ctx context.Context, req *Request) (*Response, error) { // HLxxx
	return &Response{
		Reply: fmt.Sprintf(req.GetId() + " GATEWAY-ACK @%s", time.Now().Format("2006.01.02 15:04:05")),
	}, nil
}
// END OMIT