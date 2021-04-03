package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"log"
)

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, req *wrappers.StringValue) (resp *wrappers.StringValue, err error) {
	resp = &wrappers.StringValue{}
	log.Printf("[greeter server]Hello, %s, this is greeter server.", req.Value)
	return
}
