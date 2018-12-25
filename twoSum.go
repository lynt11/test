package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		if i == len(nums)-1 {
			return []int{-1, -1}
		}
		for j := i + 1; j < len(nums); j++ {
			if res := v + nums[j]; res == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(nums, target))
}
