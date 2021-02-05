package main

import "fmt"

func main() {
END:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break END //结束至END标签
			}
			fmt.Println(i, j)
		}
	}
}
