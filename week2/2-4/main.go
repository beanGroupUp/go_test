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
	fmt.Println("%T,%V%", w, w)
	if instance, ok := w.(*Boy); ok {
		fmt.Printf("%s,%V%", instance.Name, instance)
	}
	h.GoToDinner(dest)
	fmt.Println("%T,%V%", h, h)
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

// interface{}
func SwitchInterface(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("输入为整型")
	case string:
		fmt.Println("输入为字符串")
	case Where:
		fmt.Println("实现where接口")
	}
}

func main() {
	b := Boy{Name: "小明"}
	//g := Girl{"小兰"}
	//Happy(&b, &g, "野炊吧")
	//d := Dog{Name: "大黄"}
	//Happy(&b, &d, "吃肉肉")

	SwitchInterface("Go语言极简一本通")
	SwitchInterface(99)
	SwitchInterface(&b)
}
