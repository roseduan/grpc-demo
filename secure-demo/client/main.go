package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-demo/product"
	"io/ioutil"
	"log"
)

const (
	address  = "localhost:50051"
	hostname = "localhost"

	//certFile = "secure-demo/certs/client.crt"

	clientCert = "secure-demo/certs/client.crt"
	clientKey  = "secure-demo/certs/client.key"
	caFile     = "secure-demo/certs/ca.crt"
)

func main() {
	// 单向 TLS 的代码

	//credds, err := credentials.NewClientTLSFromFile(certFile, hostname)
	//if err != nil {
	//	log.Println("new tls client err.", err)
	//	return
	//}
	//
	//opts := []grpc.DialOption{
	//	grpc.WithTransportCredentials(credds),
	//}
	//conn, err := grpc.Dial(address, opts...)
	//

	// mTLS的代码
	keyPair, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatal("failed to load client cert.", err)
		return
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal("can`t not read ca file.", err)
		return
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca file")
		return
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				ServerName:   hostname,
				Certificates: []tls.Certificate{keyPair},
				RootCAs:      certPool,
			}),
		),
	}

	conn, err := grpc.Dial(address, opts...)
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
