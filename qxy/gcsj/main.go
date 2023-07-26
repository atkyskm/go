package main

import(
	"fmt"
	"time"
)

func hello2(i int){
	println("hello goroutine:" + fmt.Sprint(i))
}

func main2(){
	for i := 5;i > 0; i--{
		go func(j int){
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}