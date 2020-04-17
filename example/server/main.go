package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/po3rin/blog-proto/rpc/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostCtl struct{}

func (p *PostCtl) One(ctx context.Context, req *post.OneReq) (*post.OneRes, error) {
	if req.GetId() == "404" {
		st := status.New(codes.NotFound, "no posts")
		return nil, st.Err()
	}

	return &post.OneRes{
		Post: &post.Post{
			Title: "this is grpc",
			Body:  "this is body",
			Tags: []string{
				"Go", "gRPC",
			},
			CreatedAt: ptypes.TimestampNow(),
			UpdatedAt: ptypes.TimestampNow(),
		},
	}, nil
}
func (p *PostCtl) List(ctx context.Context, req *post.ListReq) (*post.ListRes, error) {
	return &post.ListRes{
		Posts: []*post.Post{
			{
				Title: "this is grpc",
				Body:  "this is body",
			},
		},
	}, nil
}
func (p *PostCtl) Store(ctx context.Context, req *post.Post) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func main() {
	addr := ":50051"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	post.RegisterPostSvcServer(s, &PostCtl{})

	if err := s.Serve(l); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
