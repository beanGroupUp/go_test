package main

import "fmt"

func printFood(foods *[3]string) {
	foods[2] = "大饼"
	fmt.Println(foods)
}

func main() {
	//var array1 [6]string
	array2 := [3]string{"a", "b", "c"}
	//array3 := [...]string{"a", "b", "c"}
	//fmt.Println(array1)
	//fmt.Println(array2)
	//fmt.Println(array3)
	//
	//var matrix [3][4]string
	//var matrix2 [3][4]int
	//fmt.Println(matrix)
	//fmt.Println(matrix2)
	//
	//for i := 0; i < len(array2); i++ {
	//	fmt.Println(array2[i])
	//}
	//
	//for idx, item := range array2 {
	//	fmt.Println(idx, item)
	//}
	//
	//for _, item := range array3 {
	//	fmt.Println(item)
	//}

	printFood(&array2)
	fmt.Println(array2)
}
