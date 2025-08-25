package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []string{"葱烧海参", "烧鹅", "生蚝", "暴雨", "差天涯", "干炒牛河"}
	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go SayFoodName(v, &wg)
	}
	wg.Wait()
	fmt.Println("您点的菜,上齐了,情满用~！")
}

func SayFoodName(v string, s *sync.WaitGroup) {
	fmt.Println("您点的菜:" + v)
	s.Done()
}
