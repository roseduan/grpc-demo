package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo/product"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := product.NewProductInfoClient(conn)
	ctx := context.Background()

	id := AddProduct(ctx, client)
	GetProduct(ctx, client, id)
}

// 添加一个测试的商品
func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	aMac := &product.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success, id = ", productId.Value)
	return productId.Value
}

// 获取一个商品
func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	p, err := client.GetProduct(ctx, &product.ProductId{Value: id})
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Printf("get prodcut success : %+v\n", p)
}
