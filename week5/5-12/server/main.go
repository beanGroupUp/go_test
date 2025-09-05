package main

import (
	"fmt"
	"go_test/week5/5-12/proto/pb"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type FoodInfo struct {
	pb.UnimplementedFoodServiceServer // 嵌入未实现服务
}

// SayName 处理服务端流式 RPC 方法
func (f *FoodInfo) SayName(request *pb.FoodStreamRequest, server grpc.ServerStreamingServer[pb.FoodStreamResponse]) error {
	// 打印日志
	fmt.Println("SayName已请求")
	// 发送单个响应
	server.Send(&pb.FoodStreamResponse{Msg: "您点的菜是：" + request.Name})
	return nil
}

// PostName 处理客户端流式 RPC 方法
func (f *FoodInfo) PostName(server grpc.ClientStreamingServer[pb.FoodStreamRequest, pb.FoodStreamResponse]) error {
	// 循环接收客户端发送的流式数据
	for {
		recv, err := server.Recv() // 接收请求
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(recv.Name + ",您慢用！") // 处理请求
	}
	return nil
}

// FullStream 处理双向流式 RPC 方法
func (f *FoodInfo) FullStream(server grpc.BidiStreamingServer[pb.FoodStreamRequest, pb.FoodStreamResponse]) error {
	var wg sync.WaitGroup     // 等待组，用于同步 goroutine
	wg.Add(2)                 // 添加两个计数
	c := make(chan string, 5) // 创建缓冲通道
	// Goroutine 1: 接收客户端请求
	go func() {
		defer wg.Done() // 完成后减少等待组计数
		for {
			item, err := server.Recv() // 接收客户端消息
			if err != nil {
				fmt.Println(err)
			}
			c <- item.Name // 将消息发送到通道
			fmt.Println("已下单：" + item.Name)
			time.Sleep(time.Second * 1) // 模拟处理延迟
		}
	}()
	// Goroutine 2: 向客户端发送响应
	go func() {
		defer wg.Done() // 完成后减少等待组计数
		for {
			foodName := <-c // 从通道接收消息
			// 发送响应给客户端
			err := server.Send(&pb.FoodStreamResponse{Msg: "菜" + foodName + "做好了"})
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(time.Second * 1) // 模拟处理延迟
		}
	}()
	wg.Wait() // 等待所有 goroutine 完成

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterFoodServiceServer(server, &FoodInfo{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
