package main

import "fmt"

func main() {
	sum := 0
	i := 1
START: //标签
	if i > 100 {
		goto END
	}
	sum += i
	i++
	goto START
END:
	fmt.Println(sum)
}
