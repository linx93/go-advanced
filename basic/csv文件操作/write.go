package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

//二位数组写入csv

var data = [][]string{{"age", "name", "gender"}, {"18", "linx93", "男"}, {"18", "hwm", "女"}}

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码
	file.WriteString("\xEF\xBB\xBF")
	write := csv.NewWriter(file)

	//w.Write(data) //保存slice一维数据

	write.WriteAll(data)
	write.Flush()
	fmt.Println(data)
}
