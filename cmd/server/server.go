package main

import (
	"context"
	"fmt"
	pb "github.com/victorlin12345/grpc-playground/order"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	host = "127.0.0.1:9999"
)

type service struct {
	pb.UnimplementedOrderServer
}

var priceTable = map[string] int64{
	"YOU ARE JUST A TOY" : 2021,
}

func (s *service) CreateOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderReply, error) {
	log.Printf("Received Account %s buy %s item amount: %d", in.GetAccountID(), in.GetItemID(), in.GetAmount())
	price := priceTable[in.GetItemID()] * in.GetAmount()
	res := &pb.OrderReply{
		Status: 200,
		Message:  fmt.Sprintf("Purchase Success!, Total price %d.", price),
	}
	return res, nil
}

func main() {
	// create a http server
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on", host)
	// create a grpc server
	gRPCServer := grpc.NewServer()
	// register a service on grpc server
	pb.RegisterOrderServer(gRPCServer, &service{})
	// use service to reply the http server request
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
