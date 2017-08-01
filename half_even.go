package main

import "fmt"

func halfEven(x int) (float64, bool) {
	y := float64(x) / float64(2)
	z := x%2 == 0
	return y, z
}

func main() {
	fmt.Println(halfEven(36))
}
