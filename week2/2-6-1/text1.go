package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 指定要读取的文件名
	filename := "C:\\Users\\郑赛宇\\Desktop\\111111.txt" // 请替换为你的实际文件名

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 创建一个Scanner来读取文件
	scanner := bufio.NewScanner(file)

	// 逐行读取并打印
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%d: %s\n", lineNumber, line)
		lineNumber++
	}

	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取文件时出错: %v", err)
	}

	fmt.Println("文件读取完成")
}
