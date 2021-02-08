package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int
	var num int
	rand.Seed(time.Now().Unix())
	num = rand.Intn(100)
END:
	for t := 0; t < 4; t++ {
		fmt.Println("请猜一个数字：")
		fmt.Scan(&guess)
		if guess > num {
			fmt.Println("猜大了")
		} else if guess < num {
			fmt.Println("猜小了")

		} else {
			fmt.Println("Bingo!")
			break END
		}

	}
	fmt.Printf("正确数字是%d", num)
}
