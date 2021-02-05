package main

import (
	"fmt"
)

//行注释
/*
块注释
*/
var name string = "sy" //声明变量, var 标识符 类型
//全局变量
func main() {
	var (
		age    int = 18
		height     = 190 //省略变量类型, 从值推断变量类型
	)
	weight := 80 //短声明, 只能在函数内使用
	height = 180 //赋值
	fmt.Println(name, age, height, weight)
}
