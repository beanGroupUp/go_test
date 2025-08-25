package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

// 函数
func Show() {
	fmt.Println("从0到GO语言微服务架构师")
}

type GenerateRandom func() int

func RandomSum() GenerateRandom {
	a, b := rand.Intn(100), rand.Intn(20)
	fmt.Println(a, b)
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g GenerateRandom) Read(p []byte) (n int, err error) {
	next := g()
	if next > 2 {
		fmt.Println(">21")
		fmt.Println(next)
		fmt.Println(">21 end...")
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func PrintResult(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println("---Scan()---")
		fmt.Println(scanner.Text())
		fmt.Println("---Scan() End...---")
	}
}

type CheckOut func(int, int) int

func GetTotal(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	Show()
	show := Show
	show()
	//匿名函数
	show2 := func() {
		fmt.Println("GO语言微服务架构核心22将")
	}
	show2()

	var checkOut CheckOut = func(a int, b int) int {
		return a + b
	}
	fmt.Println(checkOut(68, 98))

	total := GetTotal(68)
	sum := total(100)
	fmt.Println(sum)
	total = GetTotal(sum)
	sum = total(50)
	fmt.Println(sum)
	total = GetTotal(sum)
	sum = total(200)
	fmt.Println(sum)

	r := RandomSum()
	PrintResult(r)
}
