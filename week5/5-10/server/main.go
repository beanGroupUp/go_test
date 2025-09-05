package main

import (
	"context"
	"fmt"
	"go_test/week5/5-9/proto/pb"
	"net"

	"google.golang.org/grpc"
)

type BookInfo struct {
}

func (b *BookInfo) Study(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{Msg: "我要学习：" + req.Name}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStudyServer(server, &BookInfo{})
	listen, err := net.Listen("tcp", "0.0.0.0:8089")
	if err != nil {
		panic(err)
	}
	fmt.Println(listen)
	err = server.Serve(listen)
	if err == nil {
		panic(err)
	}
}
