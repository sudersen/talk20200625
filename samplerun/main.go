package main

import (
	"context"
	"fmt"
	"github.com/sudersen/talk20200625/sample"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	SERVER = "localhost"
	PORT = 8080
)

func client() sample.SampleServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", SERVER, PORT), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Println("Client connected")
	return sample.NewSampleServiceClient(conn)
}

// START OMIT
func server() {
	s := grpc.NewServer()
	srv := sample.NewSampleServer()
	sample.RegisterSampleServiceServer(s, srv)
	lis, _ := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	fmt.Println("Server started")
	s.Serve(lis)
}

func main() {
	go server()
	time.Sleep(1 * time.Second)

	cli := client()
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	req := &sample.Request{Id: "123"}

	res, _ := cli.Read(ctx, req)
	fmt.Printf("message %s", res)
}
// END OMIT
