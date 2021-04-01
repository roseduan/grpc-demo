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
