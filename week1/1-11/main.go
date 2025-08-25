package main

import (
	"fmt"
)

func SoldOut(foods []string) {
	foods[1] = "已售罄"
	fmt.Println(foods)
}

func main() {
	//切片
	array5 := [...]string{"烧卤猪", "暴雨", "文昌鸡", "大龙虾", "少于"}
	//slice1 := array5[2:5]
	//fmt.Println(slice1)
	//
	//slice2 := array5[2:]
	//fmt.Println(slice2)
	//
	//slice3 := array5[:2]
	//fmt.Println(slice3)

	slice4 := array5[:]
	fmt.Println(slice4)

	//SoldOut(slice4)
	//fmt.Println(slice4)
	//
	//slice5 := slice4[1:]
	//fmt.Println(slice5)
	//
	//slice6 := slice5[2:4]
	//fmt.Println(slice6)

	//slice8 := []string{}
	//fmt.Println(slice8)
	//fmt.Printf("Address of slice8 (slice header): %p\n", &slice8)
	//slice8 = append(slice8, "蛋炒饭")
	//fmt.Println(slice8)
	//fmt.Printf("Address of slice8 (slice header): %p\n", &slice8)

	//slice8 := []string{}
	//fmt.Println(len(slice8), cap(slice8))
	//
	//for i := 0; i < 10; i++ {
	//	slice8 = append(slice8, fmt.Sprintf("蛋炒饭%d", i))
	//	fmt.Println(len(slice8), cap(slice8))
	//	fmt.Println(slice8[i])
	//
	//}
	//
	//s1 := make([]string, 10)
	//s1 = append(s1, "6666")
	//fmt.Println(len(s1))
	//s2 := make([]string, 8, 16)
	//s2 = append(s2, "go语言极简一本通")
	//fmt.Println(len(s2))
	//
	////copy
	//copy(s1, slice8)
	//copy(s2, slice8)
	//fmt.Println(s1)
	//fmt.Println(s2)

	//delete
	//slice4[1:]
	//删除最后一个
	fmt.Println(slice4[:len(slice4)-1])
	//删除第一个
	fmt.Println(slice4[1:])
}
