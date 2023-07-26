package main

import (
	"fmt"
	"math"
)

func add(a int,b int) int{ // 数据类型后置
	return a + b
}

func exists(hash map[string]int,key string) (value int,ok bool){
	value,ok = hash[key]
	return value,ok
}

func add2ptr(a *int){
	*a += 2 // 这里也要加上*
}

func main2() {
	// 变量的声明，三种:1.var，2.var + 数据类型，3. :=
	//第一种会自动识别你所赋值的数据类型
	//第二种是定义时就认为规定了他是什么数据类型
	//第三种与第一种同理，赋值时自动识别该变量的数据类型
	fmt.Println("hello,world")
	a := "hello"
	b := a + ", world"
	var c int = 1
	var d, e int = 2, 3
	var f float64
	fmt.Println(a,b,c,d,e,f)

	// 常量的声明 const

	const s string = "constant"
	const h = 50000
	const i = 3e20 / h
	fmt.Println(s,h,i,math.Sin(h),math.Sin(i))

	// if else 分支语句
	// if 后面必须跟大括号
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0{
		fmt.Println(num,"is negative")	
	}else if num < 10 {
		fmt.Println(num,"balabala")
	}else {
		fmt.Println(num,"balabala")
	}

	// 循环
	// 1.死循环
	for {
		fmt.Println("nihao")
		break
	}

	// 2.标准有限循环
	for i := 0; i < 10; i ++{
		fmt.Println(i)
	}

	// switch 分支
	nums := 4
	switch nums {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4, 5:
		fmt.Println(4,"or",5)
	default:
		fmt.Println("other")			
	}

	// 数组
	// 两种定义方式，与变量定义类似，但数组定义时必须规定数据类型
	// 1.利用var定义数组
	var arr [5] int
	arr[4] = 4
	for i := 0; i < len(arr); i ++{
		arr[i] = i
	}
	fmt.Println(arr)
	
	// 2.利用 := 来定义数组
	brr := [5] int {1,2,3,4,5}
	fmt.Println(brr)

	// 切片（变长数组）

	//利用make来创建切片
	str := make([]string, 3)
	str[0] = "a"
	str[1] = "b"
	str[2] = "c"
	fmt.Println(str,len(str))
	str = append(str,"d")
	fmt.Println(str,len(str))

	ct := make([]string, len(str))
	copy(ct,str)
	fmt.Println(ct)
	fmt.Println(ct[2:])

	// 还可以利用 := 来创建切片，就是不规定大小的数组
	bt := []string {"a","b"}
	fmt.Println(bt,len(bt))
	
	// go切片不支持负数索引

	//map 字典（哈希表）
	//map的创建用make函数来创建
	//创建map需要两个类型,括号中的是key的类型,括号外的是value的类型
	hash_map1 := make(map[string]int)
	hash_map1["atk"] = 1
	hash_map1["xjp"] = 2
	fmt.Println("hash_map1:",hash_map1)
	_,ok := hash_map1["atk"] // 查询元素是否在哈希表中,索引一个map会返回两个值，一个是其key对应的value，第二个是该key是否存在于map中
	// 如果索引值不存在，第一个返回值会返回定义value类型的零值
	_,ok1 := hash_map1["nihao"]
	fmt.Println(ok,ok1)

	//同理，map也可以直接赋值定义
	hash_map2 := map[string]int {"nihao":2}
	fmt.Println("hash_map2:",hash_map2)

	// 若要删除字典中的元素，需要用delete函数
	delete(hash_map1,"atk")
	fmt.Println("hash_map1:",hash_map1)


	// range,用来迭代数组，切片，字典或通道的关键字，且遍历map是完全无序的，完全随机~
	for key,value := range hash_map1{
		fmt.Println(key,value)
	} // 迭代字典,两个返回值分别是key和value
	// 也可以只迭代key或者只迭代value
	// for key := range hash_map1
	// for _,value := range hash_map1

	for i, item := range arr{
		fmt.Println(i,item)
	}// 迭代数组或者切片时第一个参数自动为索引，第二个参数为当前索引的值，如果不想用索引，可以用_来代替
	for _,item := range arr{
		fmt.Println(item)
	}

	// 函数
	k := 10
	l := 20
	p := add(k,l) // add 为自己定义的一个函数
	fmt.Println(p)

	val,okk := exists(hash_map1,"xjp")
	fmt.Println(val,okk)

	// 指针 go虽然有指针，但是不如c和c++一样，go的指针用途有限，主要用于传实参
	//普通函数传参一般传的都是形参，如果需要在函数内修改的话，需要传指针类型的实参
	nu := 10
	fmt.Println(nu)
	add2ptr(&nu)
	fmt.Println(nu)

}