package main

import (
	"fmt"
	"sync"
)

var (
	x    int64
	lock sync.Mutex
)

func withlock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func withoutlock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func hello(i int) {
	fmt.Println(i)
}

func main3() {
	var wg sync.WaitGroup
	src := make(chan int)
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	for i := range dest {
		fmt.Println(i)
	}

	//lock锁来保证数据在goroutine中被安全的处理
	x = 0
	wg.Add(5)
	for i := 0; i < 5; i++ {

		go func() {
			withlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("withlock", x)

	x = 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			withoutlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("withoutlock", x)

	//接下来我们我们学习WaitGroup
	//wg有三个方法
	//Add(dalta int)
	//Done()
	//Wait()
	//其中原理就是维护一个计数器，其中Add为计数器，当我们开启n个并发任务时，Add就加n，然后每个任务结束时就执行Done()，计数器就减减,最后我们用Wait来阻塞直到计数器为0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
