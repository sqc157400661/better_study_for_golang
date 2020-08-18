# 自签证书，服务加上证书验证

## 生成证书的软件
这里是[window的证书生成](http://slproweb.com/products/Win32OpenSSL.html)，其他系统自行百度

基本方法：
1. 进入bin目录
2. 执行openssl
3. 执行`genrsa -des3 -out server.key  2048`（会生成server.key，私钥文件）
4. 创建证书请求：  `req -new -key server.key -out server.csr` (会生成server.csr) 其中common.name就是域名
5. 删除密码 `rsa -in server.key -out server_no_passwd.key`
6. 执行`x509 -req -days 365 -in server.csr -signkey server_no_passwd.key -out server.crt` (会生成server.crt)

自此自签证书完成

## 服务段代码
```
package main

import (
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"services/services"
)

func main() {
	creds,err := credentials.NewServerTLSFromFile("keys/server.crt","keys/server_no_passwd.key")
	if err != nil {
		log.Fatalf("keys/server_no_passwd.key",err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}

```

## 客户端代码
```
package main

import (
	"client/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "xiaot123.com")
	if err != nil {
		log.Fatalf("加载客户端证书失败, err: %v\n", err)
	}
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
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

```