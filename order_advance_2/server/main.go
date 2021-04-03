package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	order "grpc-demo/order_advance_2"
	"log"
	"net"
)

const port = ":50051"

//初始化添加一些订单数据
func initSampleData(orderMap map[string]*order.Order) {
	orderMap["102"] = &order.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = &order.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = &order.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = &order.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = &order.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}

func main() {
	server := &OrderServer{orderMap: make(map[string]*order.Order)}
	initSampleData(server.orderMap)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	//注册订单服务
	order.RegisterOrderManagementServer(s, server)
	//注册问候服务
	order.RegisterGreeterServiceServer(s, &GreeterServer{})

	log.Println("start gRPC listen on port " + port)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

// 测试创建并传递元数据
func sendMetadata() {
	metadata.FromOutgoingContext(context.Background())
}
