package main

import (
	"fmt"
	"net/http"
)

func VisitUrl(url string) (int, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		fmt.Println("%d", res.StatusCode)
		return res.StatusCode, nil
	}
}

func SwitchShow2() (int, error) {
	if code, err := VisitUrl("http://www.baidu.com"); err != nil {
		return 0, err
	} else {
		return code, nil
	}
}

/*
*
函数式编程
*/
func ShowInfoAndPrice2(bookName string, author string, price float64) (bookInfo string, finalPrice float64) {
	bookInfo = bookName + "-作者是" + author
	finalPrice = price
	return
}

func PrintBookInfo(do func(string, string, float64) (bookInfo string, finalPrice float64),
	bookName, author string, price float64) {
	bookInfo, fianlPrice := ShowInfoAndPrice2(bookName, author, price)
	fmt.Println(bookInfo)
	fmt.Println(fianlPrice)
}

func PrintBookInfo2(bookName string, author string, price float64) {
	do := func(bookName, author string, price float64) (bookInfo string, finalPrice float64) {
		bookInfo = bookName + "作者是-" + author
		finalPrice = price
		return
	}
	//PrintBookInfo(do, bookName, author, price)
	fmt.Println(do(bookName, author, price))
}

func ShowAll(footList ...string) string {
	r := ""
	for _, item := range footList {
		r += item + "\n"
	}
	return r
}

func main() {
	fmt.Println("hello world")
	info, err := SwitchShow2()
	if err != nil {
		fmt.Println("获取百度状态：", err)
		return
	}
	fmt.Printf("获取到的百度状态码：%d\n", info)

	//PrintBookInfo(ShowInfoAndPrice2, "Go语言一本通", "欢喜", 99.00)
	//PrintBookInfo2("GO语言一本通", "欢喜", 99.00)

}
