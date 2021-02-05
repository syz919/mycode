package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入你的成绩:")
	fmt.Scan(&score)
	/*	if score >= 90 {
			fmt.Println("A")
		} else if score >= 80 {
			fmt.Println("B")
		} else if score >= 70 {
			fmt.Println("C")
		} else {
			fmt.Println("D")
		}
	*/
	switch {
	case score >= 90:
		fmt.Println("A")
		fallthrough //满足条件后继续往下进行判断
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")

	}
}
