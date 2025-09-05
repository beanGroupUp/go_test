package main

import (
	"encoding/json"
	"fmt"
	"go_test/week5/5-9/proto/pb"

	"github.com/golang/protobuf/proto"
)

type BookInfo struct {
	Name string `json:"name"`
}

func main() {
	req := pb.BookRequest{Name: "Go语言几件一本通"}
	b, err := proto.Marshal(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))

	info := BookInfo{Name: "Go语言极简一本通"}
	jsonBytes, err := json.Marshal(&info)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonBytes)
	fmt.Println(string(jsonBytes))
}
