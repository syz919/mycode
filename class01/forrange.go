package main

import "fmt"

func main() {
	a := "我爱中国"
	for v := range a {
		fmt.Println(v)
	}
	for i, v := range a { //i是索引,v是值
		fmt.Printf("%d,%q\n", i, v) //%q rune类型
	}
}
