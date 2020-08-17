module client

go 1.13

replace (
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190522204451-c2c4e71fbf69
	google.golang.org/grpc => github.com/grpc/grpc-go v1.31.0 // indirect
)

require (
	github.com/golang/protobuf v1.4.2
	golang.org/x/net v0.0.0-20200813134508-3edf25e44fcc // indirect
	golang.org/x/sys v0.0.0-20200812155832-6a926be9bd1d // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200814021100-8c09557e8a18 // indirect
	google.golang.org/grpc v1.25.1
	google.golang.org/protobuf v1.25.0
)
