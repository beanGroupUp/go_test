package main

import (
	"fmt"     // 用于格式化输出日志信息
	"net/rpc" // 提供远程过程调用（RPC）的功能
)

func main() {
	// 拨号连接到本地的 9098 端口
	c, err := rpc.Dial("tcp", "localhost:9098")
	if err != nil {
		fmt.Println(err) // 如果连接失败，输出错误并退出
		return
	}

	var reply string // 定义变量用于接收 RPC 响应
	// 调用服务端的 FoodService.SayName 方法
	// 参数 1：服务方法名 "FoodService.SayName"
	// 参数 2：传入的请求参数 "九转大肠"
	// 参数 3：用于接收响应的变量 reply 的指针
	err = c.Call("FoodService.SayName", "九转大肠", &reply)
	if err != nil {
		fmt.Println(err) // 如果调用失败，输出错误并退出
		return
	}

	// 输出服务端返回的响应
	fmt.Println(reply)
}
