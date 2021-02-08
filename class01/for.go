package main

import "fmt"

func main() {
	var num int = 4
	if num > 5 {
		fmt.Println("ok")
	} else {
		fmt.Println("nok")
	}
	fmt.Scan(&num)
	switch {
	case num > 16:
		fmt.Println("bingo")

	case num > 17:
		fmt.Println("17")
	case num > 18:
		fmt.Println("18")
	default:
		fmt.Println("stupid")
	}
}
