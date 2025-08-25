package main

import "fmt"

func main() {
	Wang := LaoWang{}
	Wang.WhereDinner()
	Xiaoli := XiaoLi{}
	Xiaoli.GoToDinner("真好")
}

type LaoWang struct {
}

type XiaoLi struct {
}

func (LaoWang) WhereDinner() {
	fmt.Println("周末吃大餐吗？")
	fmt.Println("法餐？日料？火锅？烧烤？")
}

func (XiaoLi) GoToDinner(dest string) {
	fmt.Println("女朋友：吃火锅，" + dest + "火锅店")
	fmt.Println("周末，就去" + dest + "火锅店，愉快的玩耍了")
}
