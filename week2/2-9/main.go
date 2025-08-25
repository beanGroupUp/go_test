package main

import (
	"errors"
	"fmt"
)

func funcRecover() error {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recover err:", v)
		}
	}()
	return funcCook()
}

func funcCook() error {
	panic("停水了")
	return errors.New("发生错误了")
}

func main() {
	err := funcRecover()
	if err != nil {
		fmt.Println("err is %v", err)
	} else {
		fmt.Println("err is nil")
	}
}
