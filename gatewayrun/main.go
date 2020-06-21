package main

import (
	"fmt"
	"github.com/sudersen/talk20200625/gateway"
	"github.com/sudersen/talk20200625/grpc2http"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	SERVER    = "localhost"
	GRPC_PORT = 8081
	REST_PORT = 9081
)

func client() gateway.SampleServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", SERVER, GRPC_PORT), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Println("Client connected")
	return gateway.NewSampleServiceClient(conn)
}

//cli := client()
//ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
//req := &gateway.Request{Id: "123"}
//res, _ := cli.Read(ctx, req)

// START OMIT
func server() {
	s := grpc.NewServer()
	srv := gateway.NewGateway()
	gateway.RegisterSampleServiceServer(s, srv)
	lis, _ := net.Listen("tcp", ":"+strconv.Itoa(GRPC_PORT))
	go s.Serve(lis)
	time.Sleep(1 * time.Second)
	fmt.Println("GRPC Server started on port", GRPC_PORT)

	hs := grpc2http.NewServer("0.0.0.0:"+strconv.Itoa(REST_PORT), ":"+strconv.Itoa(GRPC_PORT), // HLxxx
		gateway.RegisterSampleServiceHandlerFromEndpoint) // HLxxx
	hs.Start() // HLxxx
	fmt.Println("REST Server started on port", REST_PORT)
}

func main() {
	server()

	res, _ := http.Get("http://localhost:9081/read/123")
	bb, _ :=  ioutil.ReadAll(res.Body)
	fmt.Printf("meesage: %s\n", bb)
}
// END OMIT