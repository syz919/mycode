package main

import "fmt"

func main() {
	var name [55]string
	var score [55]int = [...]int{54: 100}
	//	var id [10]int = [10]int{1,2,3,4,6}
	var id [5]int = [5]int{3: 100, 1: 88}
	fmt.Printf("%T %T %T", name, score, id)
	fmt.Println(name, score, id)
	//数组操作
	//关系运算 == != (相同数据类型才能运算)
	var id1 [5]int = [5]int{3: 100, 1: 88}
	fmt.Println(id != id1)
	fmt.Println(id[1])
}
