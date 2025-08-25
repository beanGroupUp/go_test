package main

import "fmt"

type Run interface {
	Running()
}

type Swim interface {
	Swimming()
}

// 这是一个接口
type Sport interface {
	Run
	Swim
}

// 这是一个函数，目的是调用了 s.Running() 和 s.Swimming() 方法
func GoSports(s Sport) {
	s.Running()
	s.Swimming()
}

// 这不是一个方法，而是一个结构体类型定义
type BOY struct {
	Name string
}

// 方法定义（需要有接收者）
func (boy *BOY) Running() {
	fmt.Println(boy.Name + " is running")
}

func (boy *BOY) Swimming() {
	fmt.Println(boy.Name + " is swimming")
}

func main() {
	b := BOY{Name: "欢喜"}
	GoSports(&b)
}
