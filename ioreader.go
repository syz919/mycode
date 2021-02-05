package main

import "fmt"

func main() {
	var res string
	fmt.Print("请确认:")
	fmt.Scan(&res) //&res, 取res的指针
	fmt.Printf("你输入的是%s", res)
}
