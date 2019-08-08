package main

import (
	"fmt"
)

func main() {
	sl := make([]int, 0, 3)
	arr := [...]int{1, 2, 3}
	a := arr[0:2]
	b := arr[1:3]
	fmt.Print(sl, a, b, len(sl))
}
