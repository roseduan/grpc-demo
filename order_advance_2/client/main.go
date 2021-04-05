package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	order "grpc-demo/order_advance_2"
	"io"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()
	client := order.NewOrderManagementClient(conn)

	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.RFC3339),
		"test-key", "val1",
		"test-key", "val2",
	)
	//使用元数据context
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)
	fmt.Println("----------------use metadata----------------")
	testOrder := &order.Order{Destination: "beijing", Items: []string{"book1", "book2"}, Price: 123.232}

	//接收从服务端发送过来的metadata信息
	var header metadata.MD

	_, err = client.AddOrder(mdCtx, testOrder, grpc.Header(&header))
	log.Printf("metadata from server : %+v\n", header)

	ctx := context.Background()
	fmt.Println("----------------unary rpc----------------")
	id := AddOrder(ctx, client)
	GetOrder(ctx, client, id)

	fmt.Println("-----------server stream rpc-------------")
	SearchOrder(ctx, client)

	fmt.Println("------------client stream rpc------------")
	UpdateOrder(ctx, client)

	fmt.Println("---------bidirectional stream rpc---------")
	ProcessOrder(ctx, client)

	// 问候服务客户端
	fmt.Println("-----------greeter client rpc-----------")
	greeterClient := order.NewGreeterServiceClient(conn)
	_, err = greeterClient.SayHello(ctx, &wrappers.StringValue{Value: "roseduan"})
	if err != nil {
		log.Println("call greeter server [say hello] err.", err)
	}
}

//添加一个订单
func AddOrder(ctx context.Context, client order.OrderManagementClient) string {
	odr := &order.Order{
		Description: "a new order for test-1",
		Price:       12322.232,
		Destination: "Shanghai",
		Items:       []string{"doll", "22", "33"},
	}

	val, err := client.AddOrder(ctx, odr)
	if err != nil {
		log.Println("add order fail.", err)
		return ""
	}
	log.Println("add order success.id = ", val.String())
	return val.Value
}

//获取一个订单
func GetOrder(ctx context.Context, client order.OrderManagementClient, id string) {
	val, err := client.GetOrder(ctx, &wrappers.StringValue{Value: id})
	if err != nil {
		log.Println("get order err.", err)
		return
	}

	log.Printf("get order succes. order = %+v", val)
}

//搜索订单
func SearchOrder(ctx context.Context, client order.OrderManagementClient) {
	searchKey := "Apple"
	searchStream, _ := client.SearchOrder(ctx, &wrappers.StringValue{Value: searchKey})
	for {
		val, err := searchStream.Recv()
		if err == io.EOF { //服务端没有数据了
			break
		}
		log.Printf("search order from server : %+v", val)
	}
	return
}

//更新订单
func UpdateOrder(ctx context.Context, client order.OrderManagementClient) {
	updateStream, _ := client.UpdateOrder(ctx)
	order1 := &order.Order{Id: "103", Items: []string{"Apple Watch S6"}, Destination: "San Jose, CA", Price: 4400.00}
	order2 := &order.Order{Id: "105", Items: []string{"Amazon Kindle"}, Destination: "San Jose, CA", Price: 330.00}

	//更新订单1
	if err := updateStream.Send(order1); err != nil {
		log.Println("send order err.", err)
	}

	//更新订单2
	if err := updateStream.Send(order2); err != nil {
		log.Println("send order err.", err)
	}

	//关闭流并接收响应
	recv, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Println("close and recv err.", err)
		return
	}
	log.Printf("the update result : %+v", recv)
}

//处理订单
func ProcessOrder(ctx context.Context, client order.OrderManagementClient) {
	processStream, _ := client.ProcessOrder(ctx)

	//发送两个订单处理
	if err := processStream.Send(&wrappers.StringValue{Value: "103"}); err != nil {
		log.Println("send order err.", err)
	}
	if err := processStream.Send(&wrappers.StringValue{Value: "105"}); err != nil {
		log.Println("send order err.", err)
	}

	chn := make(chan struct{})
	//异步接收服务端的结果
	go processResultFromServer(processStream, chn)

	//再发送一个订单
	if err := processStream.Send(&wrappers.StringValue{Value: "106"}); err != nil {
		log.Println("send order err.", err)
	}
	//发送完毕后记得关闭
	if err := processStream.CloseSend(); err != nil {
		log.Println("close send err.", err)
	}

	<-chn
}

//从服务端获取处理的结果
func processResultFromServer(stream order.OrderManagement_ProcessOrderClient, chn chan struct{}) {
	defer close(chn)
	for {
		shipment, err := stream.Recv()
		if err == io.EOF {
			log.Println("[client]结束从服务端接收数据")
			break
		}
		log.Printf("[client]server process result : %+v\n", shipment)
	}
}
