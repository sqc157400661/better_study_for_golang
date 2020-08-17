package main

import (
	grpc "google.golang.org/grpc"
	"log"
	"net"
	"services/services"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
