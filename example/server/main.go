package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/po3rin/blog-proto/rpc/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostCtl struct{}

func (p *PostCtl) Get(ctx context.Context, req *post.OneReq) (*post.OneRes, error) {
	// example error ---------------------
	if req.GetId() == "404" {
		st := status.New(codes.NotFound, "no posts")
		// v := &errdetails.BadRequest{
		// 	FieldViolations: []*errdetails.BadRequest_FieldViolation{
		// 		{
		// 			Field:       "username",
		// 			Description: "should not empty",
		// 		},
		// 	},
		// }
		dt, _ := st.WithDetails()
		return nil, dt.Err()
	}

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
