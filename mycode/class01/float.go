package main

import "fmt"

func main() {
	var char string = "a\nb"  //""可解析字符串
	var char1 string = `a\nb` //''原生字符串
	var char2 string = "a\\nb"
	fmt.Println(char, char1, char2)
	name := "sy"
	name += " haha"
	fmt.Println(name)
	fmt.Printf("%T\n", name[1])
	fmt.Println(name[1])
	fmt.Println(len(name))
	fmt.Println(name[0:4])
	//age := "y"
	//fmt.Println(uint8(age))
}
