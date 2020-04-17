package main

import (
	"context"
	"fmt"
	"os"

	"github.com/po3rin/blog-proto/rpc/post"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	res, err := post.NewPostSvcClient(conn).One(context.Background(), &post.OneReq{Id: "hello grpc"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", res.GetPost())
}
