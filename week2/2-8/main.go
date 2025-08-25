package main

import (
	"bufio"
	"fmt"
	"os"
)

func Cook() {
	//栈
	defer fmt.Println("吃点好的，开饭")
	defer fmt.Println("播放音乐")
	fmt.Println("买菜")
	fmt.Println("卖肉")
	panic("停水了，停气了")
	fmt.Println("做饭")
}

func WriteMenu(filename string, foods []string) {
	curDir, _ := os.Getwd()
	f, err := os.Create(curDir + "/" + filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, food := range foods {
		fmt.Fprintln(w, food)
	}
}

func main() {
	//加了defer后，会到最后才执行
	//Cook()
	s := []string{"葱烧海参", "烧鹅", "谈考生号", "暴雨", "差天涯", "干炒牛河"}
	WriteMenu("/week2/2-8/menu.txt", s)
}
