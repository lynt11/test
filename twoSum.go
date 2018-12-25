package main

import "fmt"

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i, v := range nums {
		maps[v] = i
	}
	for i, v := range nums {
		temp := target - v
		if _, isok := maps[temp]; isok && maps[temp] != i {
			return []int{i, maps[temp]}
		}
	}
	return []int{-1, -1}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(nums, target))
}
