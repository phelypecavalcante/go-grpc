deps:
	$ brew install protobuf
	$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

compile-protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
  	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
  	proto/hello.proto

server: 
	go run cmd/server/main.go
	
client:
	go run cmd/client/main.go