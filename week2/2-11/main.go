package main

import (
	"fmt"
	"time"
)

func ShowBook() {
	fmt.Println("欢喜-Go语言极简一本通")
}

func main() {
	//go
	go ShowBook()

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(fmt.Sprintf("i am %d", i))
		}()
	}

	time.Sleep(1 * time.Second)
}
