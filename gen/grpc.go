package gen

//go:generate protoc -I=../protos --go_out=../pkg/pb --go_opt=paths=source_relative --go-grpc_out=../pkg/pb --go-grpc_opt=paths=source_relative base_sdk.proto request.proto response.proto
