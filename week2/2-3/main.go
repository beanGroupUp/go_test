package main

import "fmt"

// 接口的实现
type Boy struct {
	Name string
}

type Where interface {
	WhereDinner()
}

type GoToHappy interface {
	GoToDinner(dest string)
}

func Happy(w Where, h GoToHappy, dest string) {
	w.WhereDinner()
	h.GoToDinner(dest)
}

func (b *Boy) WhereDinner() {
	fmt.Println(b.Name + ",周末去吃大餐吗？")
}

type Girl struct {
	Name string
}

func (g *Girl) GoToDinner(dest string) {
	fmt.Println(g.Name + "去" + dest)
}

type Dog struct {
	Name string
}

func (d *Dog) GoToDinner(dest string) {
	fmt.Println(d.Name + "在家" + dest)
}

func main() {
	b := Boy{Name: "小明"}
	g := Girl{"小兰"}
	Happy(&b, &g, "野炊吧")
	d := Dog{Name: "大黄"}
	Happy(&b, &d, "吃肉肉")
}
