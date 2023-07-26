package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main4(){
	rand.Seed(time.Now().Unix())
	n := rand.Intn(101)
	for {
		fmt.Print("请输入您要猜测的数字:")
		reader := bufio.NewReader(os.Stdin)
		input,err := reader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
		}
		input = strings.TrimSuffix(input,"\r\n")
		guess,err := strconv.Atoi(input)
		fmt.Print(guess)
		if guess > n{
			fmt.Println("大了")
		}else if guess < n{
			fmt.Println("小了")
		}else{
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}