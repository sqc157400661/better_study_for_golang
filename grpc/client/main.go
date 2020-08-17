package main

import (
	"client/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//Set up a connection to the server
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	prodClient := services.NewProductServiceClient(conn)
	prodRes, err := prodClient.GetProductStock(context.Background(),
		&services.ProdRequest{ProdId: 12})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes.ProdStock)
}
