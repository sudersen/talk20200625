package sample

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	SERVER = "localhost"
	PORT = 8000
)

func Read() { // HLxxx
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", SERVER, PORT), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	cli := NewSampleServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	req := &Request{Id: "123"}

	res, err := cli.Read(ctx, req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("message %s", res)
}