package main

import (
	"fmt"
	"time"
)

func main() {

	/*	stat := time.Now()
		ch1 := make(chan interface{})
		ch2 := make(chan string)
		ch3 := make(chan string)

		go func() {
			time.Sleep(4 * time.Second)
			close(ch1)
		}()

		go func() {
			time.Sleep(3 * time.Second)
			ch2 <- "从0到Go语言微服务架构师"
		}()

		go func() {
			time.Sleep(3 * time.Second)
			ch3 <- "Go语言微服务架构核心22讲"
		}()

		fmt.Println("等待读取....")

		select {
		case <-ch1:
			fmt.Printf("未阻塞%v\n", time.Since(stat))
		case data := <-ch2:
			fmt.Printf("ch2 %v", data)
		case data := <-ch3:
			fmt.Printf("ch3 %v", data)
				//default:
				//fmt.Printf("默认值")
		}*/

	//执行顺序：从上到下，从左到右
	/*	ch4 := make(chan string, 8)
		ch5 := make(chan string, 8)
		chs := []chan string{ch4, ch5}
		names := []string{"欢喜", "Go语言极简一本通", "面向加薪学习"}
		select {
		case getChan(0, chs) <- getName(2, names):
			fmt.Println("执行了第一条语句")
		case getChan(1, chs) <- getName(1, names):
			fmt.Println("执行了第二条语句")
		default:
			fmt.Println("默认...")
		}*/

	ch := make(chan string)
	go func() {
		for {
			ch <- "烤乳猪"
		}
	}()

	for {
		select {
		case data := <-ch:
			fmt.Println(data)
			break
		}
		goto exit
		time.Sleep(3 * time.Second)
		fmt.Println("循环中...")
	}

exit:
	fmt.Println("跳出循环")
	fmt.Println("结束")
}

func getName(i int, names []string) string {
	fmt.Println("names:%d\n", i)
	return names[i]
}

func getChan(i int, chs []chan string) chan string {
	fmt.Println("chs:%d\n", i)
	return chs[i]
}
