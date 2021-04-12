package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"grpc-demo/load_balance_demo"
	"log"
	"time"
)

var addrs = []string{"localhost:50051", "localhost:50052"}

const (
	exampleScheme      = "example"
	exampleServiceName = "lb.example.com"
)

func main() {
	conn, _ := grpc.Dial(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		grpc.WithBalancerName(grpc.PickFirstBalancerName),
		grpc.WithInsecure(),
	)
	defer conn.Close()

	makeRPCs(conn, 10)
}

func makeRPCs(cc *grpc.ClientConn, n int) {
	client := load_balance_demo.NewEchoServiceClient(cc)
	for i := 0; i < n; i++ {
		callUnaryEcho(client, "test for load balance")
	}
}

func callUnaryEcho(c load_balance_demo.EchoServiceClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &wrappers.StringValue{Value: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Value)
}

type exampleResolverBuilder struct{}

type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: addrs,
		},
	}

	r.start()
	return r, nil
}

func (*exampleResolverBuilder) Scheme() string { return exampleScheme }

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}

func init() {
	resolver.Register(&exampleResolverBuilder{})
}
