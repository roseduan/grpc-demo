package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// 一元拦截器
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (res interface{}, err error) {
	//前置处理
	log.Println("==========[Server Unary Interceptor]===========", info.FullMethod)

	//完成方法的正常执行
	res, err = handler(ctx, req)

	//后置处理
	log.Printf("After method call, res = %+v\n", res)
	return
}

// 服务端流拦截器
type WrappedServerStream struct {
	grpc.ServerStream
}

func (w *WrappedServerStream) SendMsg(m interface{}) error {
	log.Printf("[order stream server interceptor] send a msg : %+v", m)
	return w.ServerStream.SendMsg(m)
}

func (w *WrappedServerStream) RecvMsg(m interface{}) error {
	log.Printf("[order stream server interceptor] recv a msg : %+v", m)
	return w.ServerStream.RecvMsg(m)
}

func NewWrappedServerStream(s grpc.ServerStream) *WrappedServerStream {
	return &WrappedServerStream{s}
}

func orderStreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("=========[order stream]start %s\n", info.FullMethod)

	//执行方法
	err := handler(srv, NewWrappedServerStream(ss))
	if err != nil {
		log.Println("handle method err.", err)
	}

	log.Printf("=========[order stream]end")
	return nil
}
