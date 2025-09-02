package main

import (
	"fmt"
	"go_test/week1/1-14/hello"
	model2 "go_test/week1/1-14/hello2"
)

func main() {
	var Food model.Food
	var Player model2.Player
	fmt.Println(Food)
	fmt.Println(Player)

	price := 68
	newPrice := model.MyInt(price)
	fmt.Println(newPrice)

	food := model.Food{Name: "豆汁儿"}
	p := model.Person{Food: food}
	p.SayHi("老张")

	food2 := model.Food2{
		FoodName:   "葱烧海参",
		FoodSystem: model.FoodSystem{Name: "鲁菜"},
		Another:    model.FoodSystem{Name: "豫菜"},
	}
	fmt.Println(food2.Another.Name)
	fmt.Println(food2.Name)
	fmt.Println(food2.FoodSystem.Name)
}
