package main

import "fmt"

func findMax(nums ...float64) float64 {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(findMax(2, 36, 44, 419))
}
