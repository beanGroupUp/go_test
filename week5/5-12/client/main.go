package main

import (
	"context"
	"fmt"
	"go_test/week5/5-12/proto/pb"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	// 连接到 gRPC 服务器（禁用安全传输）
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()                      // 确保连接关闭
	client := pb.NewFoodServiceClient(conn) // 创建食品服务客户端
	//服务端流模式
	//res, err := client.SayName(context.Background(), &pb.FoodStreamRequest{Name: "干锅肥肠"})
	//if err != nil {
	//	panic(err)
	//}
	//// 循环接收服务端的流式响应
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv.Msg) // 打印响应消息
	//}

	////客户端流模式
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancelFunc() // 确保取消上下文
	//postNameClient, err := client.PostName(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//foods := []string{"东坡肘子", "回锅肉", "夫妻肺片", "鱼香茄子", "辣子鸡丁", "水煮牛肉", "宫保鸡丁"}
	//// 循环发送多个请求
	//for _, food := range foods {
	//	fmt.Println("上菜：" + food)
	//	err := postNameClient.Send(&pb.FoodStreamRequest{Name: food})
	//	time.Sleep(time.Second * 1) // 模拟延迟
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//}

	//双向流模式
	foodsDouble := []string{"东坡肘子", "回锅肉", "夫妻肺片", "鱼香茄子", "辣子鸡丁", "水煮牛肉", "宫保鸡丁"}
	fullClient, err := client.FullStream(context.Background())
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup // 等待组，用于同步 goroutine
	wg.Add(2)             // 添加两个计数
	// Goroutine 1: 接收服务端的流式响应
	go func() {
		defer wg.Done() // 完成后减少等待组计数
		for {
			item, err2 := fullClient.Recv()
			if err2 != nil {
				fmt.Println(err2)
			}
			fmt.Println(item.Msg) // 打印响应消息
		}
	}()
	// Goroutine 2: 向服务端发送流式请求
	go func(s []string) {
		defer wg.Done() // 完成后减少等待组计数
		for _, item := range foodsDouble {
			err2 := fullClient.Send(&pb.FoodStreamRequest{Name: item})
			if err2 != nil {
				fmt.Println(err2)
			}
		}
	}(foodsDouble)

	wg.Wait() // 等待所有 goroutine 完成
}
