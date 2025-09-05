package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Msg string `json:"msg"`
}

func GetUrl(url string) (string, error) {
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var result Result
	err = json.Unmarshal(b, &result)
	if err != nil {
		return "", err
	}
	return result.Msg, err
}

func main() {
	//go语言自带的http调用
	/*	r, err := http.Get("http://127.0.0.1:9098/hello?lesson=从0到Go语言微服务架构师")
		//关闭
		defer r.Body.Close()
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var result Result
		err = json.Unmarshal(b, &result)
		if err != nil {
			panic(err)
		}
		fmt.Println(".......................")
		fmt.Println(result.Msg)*/

	url := "http://127.0.0.1:9098/hello?lesson=从0到Go语言微服务架构师"
	result, err := GetUrl(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("调用结果：", result)
}
