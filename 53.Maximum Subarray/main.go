package main

import (
	"fmt"
)

func main() {
	result := maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		t := sum + nums[i]
		result = max(result, t)
		if t <= 0 {
			sum = 0
		} else {
			sum = t
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
