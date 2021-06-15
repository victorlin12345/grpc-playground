package main

import (
	"context"
	pb "github.com/victorlin12345/grpc-playground/order"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:9999"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := pb.OrderRequest{
		AccountID: "victorlin54321",
		ItemID: "YOU ARE JUST A TOY",
		Amount: 10,
	}
	res, err := c.CreateOrder(ctx, &req)
	if err != nil {
		log.Fatalf("could not order: %v", err)
	}
	log.Printf("Request Status: %d", res.Status)
	log.Printf("Response Message: %s", res.Message)

}