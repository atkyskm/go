package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type point struct{
	name,password int
}

type user struct{
	Name string
	Age int
}

func main3() {
	a := "hellllo"
	fmt.Println(strings.Contains(a,"ll")) // 字符串匹配，查看a是否包含第二个参数的字符串
	fmt.Println(strings.Count(a,"ll")) // 字符串计数，计算a中包含几个ll
	fmt.Println(strings.Index(a,"e")) // 查找字符串的索引下标
	fmt.Println(strings.Join([]string{"hello","world"},"-")) // 将数组中的字符串连接起来，第二个参数为连接项
	fmt.Println(strings.Repeat(a,2)) // 将字符串复制几次
	fmt.Println(strings.Replace(a,"e","E",-1)) // 将字符串中的目标字符替换为我们规定的字符，最后一个参数为替换几次，-1代表所有全部替换
	fmt.Println(strings.Split("he-ll-o","-")) // 与python的split函数一样，将连接符删掉并返回断掉后的数组
	fmt.Println(strings.ToLower(a)) // 将字符串所有字符转换为小写
	fmt.Println(strings.ToUpper(a)) // 将字符串所有字符转换为大写
	b := strings.Split(a,"") // 通过这种形式可以将字符串转化为字符数组
	fmt.Println(b)

	// 字符串格式化 就是c里面的prinf,数字%d，字符串%s，字符%c，%f为浮点数，%v为万能类型，会自动识别
	p := point{1,2}
	fmt.Printf("p=%+v\n",p)

	c := user{"atk",20}
	buf,err := json.Marshal(c) // json转码只识别大写字母开头的结构体变量
	fmt.Println(buf,string(buf),err)
	var sp user
	err = json.Unmarshal(buf,&sp)
	fmt.Println(sp)
}