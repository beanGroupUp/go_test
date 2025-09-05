package main

import (
	"context"
	"fmt"
	"go_test/week5/5-9/proto/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewStudyClient(conn)
	res, err := client.Study(context.Background(), &pb.BookRequest{Name: "从0到Go语言微服务架构师"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Msg)
}
