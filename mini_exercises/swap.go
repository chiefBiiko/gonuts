package main

import "fmt"

func swap(a *int, b *int) {
  tmp := *a
  *a = *b
  *b = tmp
}

func main() {
  x, y := 419, 401
  fmt.Println("x:", x, "y:", y)
  swap(&x, &y)
  fmt.Println("x:", x, "y:", y)
}