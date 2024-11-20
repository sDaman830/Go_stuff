package main

import (
	"fmt"
)

func main() {
	var thingy [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		fmt.Println(thingy[i])
	}
}
