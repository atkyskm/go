package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main3(){
	rand.Seed(time.Now().Unix())
	n := rand.Intn(101)
	for {
		var num int
		println("请输入您要猜的数字")
		fmt.Scanln(&num)
		if num < n{
			println("小了")
		}else if num > n{
			println("大了")
		}else{
			println("恭喜你，猜对了")
			break
		}
	}
}