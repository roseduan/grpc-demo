package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	order "grpc-demo/order_advance_2"
	"log"
	"net"
)

var addrs = []string{
	"50051",
	"50052",
}

type EchoServer struct {
	addr string
}

func (e *EchoServer) UnaryEcho(ctx context.Context, req *order.EchoReq) (resp *order.EchoResp, err error) {
	return &order.EchoResp{
		Message: fmt.Sprintf("%s from(%s)", req.Message, e.addr),
		Addr:    e.addr,
	}, nil
}

func startServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("net listen err.", err)
		return
	}

	s := grpc.NewServer()
	order.RegisterEchoServiceServer(s, &EchoServer{addr})
	log.Printf("serve on %s", addr)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve %+v", err)
	}
}
