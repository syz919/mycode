package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var isbody bool
	/*
		布尔类型, 0值是false, 字面量true/false
	*/
	fmt.Printf("%T,%t\n", isbody, isbody) //%t, 布尔类型值, %T, 数据类型
	fmt.Println(unsafe.Sizeof(isbody))    //计算标识符长度

	var num int = 16
	var char byte = 'a'
	var codepoint rune = '我'
	fmt.Printf("%T\n", char)
	fmt.Printf("%T\n", codepoint)
	fmt.Println(char, codepoint)
	fmt.Printf("%d,%c,%U\n", num, char, codepoint) //%d 10进制整数  %c byte %U unicode
	fmt.Printf("%b\n", num)                        //%b 二进制
	/*
		位运算  & | ^ << >> &^
	*/
	a := 10
	var b int8 = 8
	fmt.Println(a + int(b))
	c := 128
	fmt.Println(int8(c))
	fmt.Printf("%10d %-10d %010d %-010d %d\n", c, c, c, c, c)
}
