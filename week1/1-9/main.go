package main

import "fmt"

func byValue(price int) {
	//此时的运算值作用在方法内部，出了方法参数值不变
	price += 20
	fmt.Println("price is", price)
}

func byReference(price *int) {
	//作用在内存中
	*price += 20
	fmt.Println("price is", *price)
}

func main() {
	var price int = 68
	//var ptr *int = &price
	//fmt.Println(ptr)
	//
	//*ptr = 8
	//fmt.Println(price)

	//值拷贝
	byValue(price)
	fmt.Println(price)
	//引用拷贝
	byReference(&price)
	fmt.Println(price)
}
