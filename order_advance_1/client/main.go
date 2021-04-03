package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"grpc-demo/order"
	"io"
	"log"
	"sync"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(UnaryClientOrderInterceptor), //注册拦截器
		grpc.WithStreamInterceptor(StreamClientOrderInterceptor),
	)
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	// 使用带有截止时间的context
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second)) //适当调整截止时间观察不同的调用效果
	defer cancel()

	client := order.NewOrderManagementClient(conn)
	AddOrder(ctx, client)

	//fmt.Println("----------------unary rpc----------------")
	//id := AddOrder(ctx, client)
	//GetOrder(ctx, client, id)

	fmt.Println("----------request with deadline----------")
	chn := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer func() {
			fmt.Println("添加订单成功")
			wg.Done()
		}()
		id := AddOrder(ctx, client)
		time.Sleep(1 * time.Second)
		chn <- id
	}()
	go func() {
		id := <-chn
		time.Sleep(2 * time.Second)
		GetOrder(ctx, client, id)
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("-----------cancel rpc request------------")
	cancelRpcRequest(client)

	ctx = context.Background()
	fmt.Println("-----------server stream rpc-------------")
	SearchOrder(ctx, client)

	fmt.Println("------------client stream rpc------------")
	UpdateOrder(ctx, client)

	fmt.Println("---------bidirectional stream rpc---------")
	ProcessOrder(ctx, client)
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

// 取消RPC请求
func cancelRpcRequest(client order.OrderManagementClient) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	done := make(chan string)
	go func() {
		var id string
		defer func() {
			fmt.Println("结束执行, id = ", id)
			done <- id
		}()

		time.Sleep(2 * time.Second)
		id = AddOrder(ctx, client)
		log.Println("添加订单成功, id = ", id)
	}()

	//等待一秒后取消
	time.Sleep(time.Second)
	cancelFunc()

	<-done
}
