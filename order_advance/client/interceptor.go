package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

//客户端一元拦截器
func UnaryClientOrderInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	log.Println("=========[client interceptor] ", method)
	err = invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		log.Println("invoke method err.", err)
	}
	log.Println("=========[client interceptor] end. reply : ", reply)
	return
}

// 客户端流拦截器
type WrappedClientStream struct {
	grpc.ClientStream
}

func (w *WrappedClientStream) SendMsg(m interface{}) error {
	log.Printf("===========[client interceptor] send msg : %+v", m)
	return w.ClientStream.SendMsg(m)
}

func (w *WrappedClientStream) RecvMsg(m interface{}) error {
	log.Printf("============[client interceptor] recv msg : %+v", m)
	return w.ClientStream.RecvMsg(m)
}

func NewWrappedClientStream(s grpc.ClientStream) *WrappedClientStream {
	return &WrappedClientStream{s}
}

func StreamClientOrderInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Printf("===========[client msg]start, method = %+v\n", method)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return NewWrappedClientStream(clientStream), nil
}
