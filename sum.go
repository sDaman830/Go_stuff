package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {
	a := 10
	b := 20

	if b > 18 {
		b = 18
	} else {

	}
	fmt.Println(add(a, b))
}
