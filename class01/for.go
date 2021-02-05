package main

import "fmt"

func main() {
	/*
		for x;y;z 初始化子语句;条件子语句;后置子语句(循环体执行完之后的语句){循环体
		}
	*/
	sum := 0
	/*	for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
