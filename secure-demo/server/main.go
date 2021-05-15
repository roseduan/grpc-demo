package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"grpc-demo/product"
	"io/ioutil"
	"log"
	"net"
)

const (
	port = ":50051"

	//certFile = "secure-demo/certs/server.crt"
	//keyFile = "secure-demo/certs/server.key"

	serverCert = "secure-demo/certs/server.crt"
	serverKey  = "secure-demo/certs/server.key"
	caFile     = "secure-demo/certs/ca.crt"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	// 单向 TLS 的代码
	//// 加载证书
	//x509KeyPair, err := tls.LoadX509KeyPair(certFile, keyFile)
	//if err != nil {
	//	log.Println("load cert err.", err)
	//	return
	//}
	//
	//opts := []grpc.ServerOption{
	//	grpc.Creds(credentials.NewServerTLSFromCert(&x509KeyPair)),
	//}
	//
	//s := grpc.NewServer(opts...)

	// mTLS 的代码
	keyPair, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		log.Fatal("failed to load key pair.", err)
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

	// 为所有传入的连接启用TLS
	opts := []grpc.ServerOption{
		grpc.Creds(
			credentials.NewTLS(&tls.Config{
				ClientAuth:   tls.RequestClientCert,
				Certificates: []tls.Certificate{keyPair},
				ClientCAs:    certPool,
			}),
		),
	}

	s := grpc.NewServer(opts...)
	product.RegisterProductInfoServer(s, &server{})
	log.Println("start gRPC listen on port " + port)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

type server struct {
	productMap map[string]*product.Product
}

//添加商品
func (s *server) AddProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

//获取商品
func (s *server) GetProduct(ctx context.Context, req *product.ProductId) (resp *product.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	resp = s.productMap[req.Value]
	return
}
