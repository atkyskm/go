package main

import (
	"fmt"
	"strconv"
	"time"
)

func main5(){
	now := time.Now()
	fmt.Println(now)
	shijianchuo := now.Unix()
	fmt.Println(shijianchuo)

	f, _ := strconv.ParseFloat("1.234",64) // 将字符串转换为数字，第一个参数为待转换的字符串，第二个参数为转换的进制数，不写代表自动识别
	fmt.Println(f)

	// 也可以用Atoi来自动识别并转换为数字
	n, _ := strconv.Atoi("123.5") // Atoi只能转化整数
	fmt.Println(n)

	n1, _ := strconv.Atoi("111")
	fmt.Println(n1)

	n2 := strconv.Itoa(123)
	fmt.Println(n2)
}