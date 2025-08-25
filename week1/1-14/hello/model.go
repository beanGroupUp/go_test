package model

import "fmt"

func Cook(f Food) {}

func Send(f Food) {}

type MyInt int32

type hello func(string)

type Person struct {
	Food Food
}

type Food struct {
	Name string
}

func (p *Person) SayHi(name string) {
	fmt.Println(name + "吃了吗？" + p.Food.Name)
}

type FoodSystem struct {
	Name string
}

type Food2 struct {
	FoodName string
	FoodSystem
	Another FoodSystem
}
