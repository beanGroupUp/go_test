package main

import (
	"fmt"
	"net"
	"net/rpc"
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
	// 监听本地 9098 端口，等待客户端连接
	listen, err := net.Listen("tcp", ":9098")
	if err != nil {
		fmt.Println(err) // 如果监听失败，输出错误并退出
		return
	}

	/**
	rpc.RegisterName 是 Go 标准库 net/rpc 提供的函数，用于将一个本地服务（&FoodService{}）注册到 RPC 服务器中，使其可以通过指定的名称（"FoodService"）被远程调用。
	"FoodService" 是服务的唯一标识符，客户端调用时会用到这个名称。
	&FoodService{} 是服务的实例，必须是一个指针类型（因为 RPC 要求方法接收者是指针类型）。




	2. 服务注册的条件
	要成功注册服务，FoodService 必须满足以下条件：

	方法签名要求：
	方法必须有两个参数，且第一个参数是输入参数，第二个参数是指向输出参数的指针。
	方法必须返回 error 类型。
	示例：
	go
	深色版本
	func (s *FoodService) GetMenu(request string, reply *string) error {
	    *reply = "Today's menu: Pizza"
	    return nil
	}
	方法可见性：
	方法必须是 公开的（首字母大写），否则不会被注册。
	结构体定义：
	FoodService 结构体需要定义符合上述规则的方法。
	*/
	// 注册 RPC 服务，将 FoodService 注册为名为 "FoodService" 的服务
	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		fmt.Println(err) // 如果注册失败，输出错误并退出
		return
	}

	// 阻塞等待客户端连接
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(err) // 如果连接失败，输出错误并退出
		return
	}

	// 使用 RPC 协议处理客户端连接，调用 SayName 方法
	// 注意：此代码仅处理一个连接，处理完后程序结束
	rpc.ServeConn(conn)
}
