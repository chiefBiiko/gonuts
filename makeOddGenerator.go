package main

import "fmt"

func makeOddGenerator() func() uint {
  odd := uint(1)
  return func() uint {
    rtn := odd
    odd += 2
    return rtn
  }
}

func main() {
  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
}