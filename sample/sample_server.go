package sample

import (
	"context"
	"fmt"
	"time"
)

// START OMIT

func NewSampleServer() *sampleServer { // HLxxx
	return &sampleServer{}
}

type sampleServer struct{} // HLxxx

func (s *sampleServer) Read(ctx context.Context, req *Request) (*Response, error) { // HLxxx
	return &Response{
		Reply: fmt.Sprintf(req.GetId() + " ACK @%s", time.Now().Format("2006.01.02 15:04:05")),
	}, nil
}
// END OMIT