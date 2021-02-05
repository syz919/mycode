package main

import (
	"fmt"
)

const (
	name      string = "哈哈"
	statusnew int    = 1
)

func main() {
	const (
		a = 1
		b        //在一个小括号内如果没有赋值, 则使用最近的已赋值常量对应的值进行赋值
		c = iota //枚举, 通常用于常量, 在一个小括号内, 初始值为0, 每调用一次+1
		d
		e
		f
	)
	const (
		aa = iota
		bb
		cc
		dd
	)
	fmt.Println(name, statusnew)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(aa, bb, cc, dd)

	{
		var scale0 = 30
		fmt.Println(scale0)
	}

	var scale0 = 40
	fmt.Println(scale0)

	var ns string = "sy"

	fmt.Print("1.haha")
	fmt.Print("2.xixi")    //print没有换行符
	fmt.Printf("%s\n", ns) //printf, 打印占位符

}
