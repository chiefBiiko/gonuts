package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() (rtn int) {
		rtn, a, b = a, b, a + b
		return
	}
}

func main() {
  fibo := fibonacci()
	for i := 0; i < 10; i++ {
	  fmt.Println(fibo())
	}
}
