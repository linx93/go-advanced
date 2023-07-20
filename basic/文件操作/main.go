package main

import (
	"fmt"
	"os"
)

func main() {
	//read()
	//write()
}

func read() ([]byte, error) {

	//file, err := os.Open("D:\\code\\go\\linx\\go-advanced\\basic\\文件操作\\a.txt")
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return nil, err
	}

	defer file.Close()

	content := make([]byte, 100)

	n, err := file.Read(content)
	if err != nil {
		fmt.Println("读文件发生错误", err)
		return nil, err
	}
	bytes := content[:n]

	fmt.Println(string(bytes))

	return content, nil
}

func write() {
	file, err := os.OpenFile("b.txt", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("创建文件错误", err)
		return
	}
	defer file.Close()
}
