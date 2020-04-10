package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/po3rin/blog-proto/rpc/post"
	"google.golang.org/grpc"
)

type PostCtl struct{}

func (p *PostCtl) Get(context.Context, *post.OneReq) (*post.OneRes, error) {
	return &post.OneRes{
		Title: "this is grpc",
		Body:  "this is body",
	}, nil
}

func main() {
	addr := ":50051"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	post.RegisterPostServer(s, &PostCtl{})

	if err := s.Serve(l); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
