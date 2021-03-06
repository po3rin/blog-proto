gen:
	protoc --proto_path proto/ post/post.proto --go_out=plugins=grpc:rpc/

# example ----------------------

server:
	GODEBUG=http2debug=2 go run example/server/main.go

client:
	GODEBUG=http2debug=2 go run example/client/main.go
