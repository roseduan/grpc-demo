package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"grpc-demo/load_balance_demo"
	"log"
	"net"
	"sync"
)

var addrs = []string{":50051", ":50052"}

type Server struct {
	addr string
}

func (s *Server) SayHello(ctx context.Context, req *wrappers.StringValue) (resp *wrappers.StringValue, err error) {
	resp = &wrappers.StringValue{}
	log.Println("the server port is ", s.addr)
	return
}

func startServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("tcp listen err.", err)
		return
	}
	s := grpc.NewServer()
	load_balance_demo.RegisterEchoServiceServer(s, &Server{addr})
	log.Printf("serving on %s\n", addr)
	if err := s.Serve(listener); err == nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			startServer(val)
		}(addr)
	}
	wg.Wait()
}
