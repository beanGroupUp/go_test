package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8089")
	if err != nil {
		fmt.Println(err)
		return
	}
	reply := ""

	//xxx.MethodName
	//json

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))
	//{“method”:"FoodService.SayName",params:["爆炒小龙虾"],"id":1}
	err = client.Call("FoodService.SayName", "丰田就吃", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)
}
