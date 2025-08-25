package main

import (
	"fmt"
	"sort"
)

func main() {
	//三种创建map的方式
	m := make(map[string]int)
	fmt.Println(m)

	var m2 map[string]int
	fmt.Println(m2)

	ma3 := map[string]string{
		"鲁菜": "九转大肠",
		"川菜": "麻婆豆腐",
	}
	fmt.Println(ma3)

	m4 := map[string]map[int]string{}
	fmt.Println(m4)

	m5 := map[int]string{}
	m5[1] = "葱烧海参"
	m5[2] = "回锅肉"
	for k, v := range m5 {
		fmt.Println(k, v)
	}

	for _, v := range m5 {
		fmt.Println(v)
	}

	keys := make([]int, 0, len(m5))
	for k, _ := range m5 {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, m5[k])
	}

	v := m5[1]
	fmt.Println(v)

	v2, ok := m5[10]
	if ok {
		fmt.Println(v2)
	} else {
		fmt.Println("找不到数据~！")
	}

	delete(m5, 1)
	fmt.Println(m5)
	//删除不存在的key不会报错
	delete(m5, 10)
	fmt.Println(m5)
}
