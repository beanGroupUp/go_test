package main

import "fmt"

func GivenFood() chan string {
	ch := make(chan string)
	go func() {
		ch <- "回锅肉"
		ch <- "碳烤生蚝"
		ch <- "担担面"
		close(ch)
	}()
	return ch
}

// 只接收
func OnlyReceive(ch chan<- string) {
	ch <- "hello"
	close(ch)
}

// 只读
func OnlyRead(ch <-chan string) {
	//第二种方式
	for data := range ch {
		fmt.Println(data)
	}
}

func main() {
	//channel

	//ch := make(chan string)
	////缓冲通道，可缓冲的容量为6
	//ch2 := make(chan string, 6)
	//ch <- "回锅肉"
	//<-ch2
	//close(ch)

	ch := make(chan string)
	/*	ch = GivenFood()
		//第一种方式
		for {
			if name, ok := <-ch; ok {
				fmt.Println(name)
			} else {
				break
			}
		}

		//第二种方式
		for data := range ch {
			fmt.Println(data)
		}*/

	go OnlyReceive(ch)
	OnlyRead(ch)

}
