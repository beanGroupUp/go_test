package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义一个 FoodService 结构体，作为 RPC 服务的载体
type FoodService struct {
}

// 实现 RPC 方法 SayName
// request 是客户端传入的参数，resp 是返回给客户端的响应
func (f *FoodService) SayName(request string, resp *string) error {
	*resp = "您点的菜是：" + request // 拼接响应字符串
	return nil                 // 返回 nil 表示没有错误
}

func main() {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		//rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
