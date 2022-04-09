package main

import (
	"fmt"
	"math"
)

func main() {
	result := maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	lowestNegativeNumber := 0 - math.MaxInt
	temp := map[int]bool{}
	sum := 0

	for _, val := range nums {
		if temp[val] != true && val > 0 {
			sum += val
			temp[val] = true
		}

		if val < 0 && val > lowestNegativeNumber {
			lowestNegativeNumber = val
		}
	}

	if sum == 0 {
		return 0
	}

	fmt.Println(sum)
	fmt.Println(lowestNegativeNumber)
	return sum + lowestNegativeNumber
}
