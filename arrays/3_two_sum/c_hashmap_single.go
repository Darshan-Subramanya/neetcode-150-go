package main

import (
	"fmt"
)

func twoSumHashMapSinglePass(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, val := range nums {
		need := target - val
		j, exists := seen[need]
		if exists {
			return []int{j, i}
		}
		seen[val] = i
	}
	return nil
}

func main() {
	nums := []int{6, 7, 3, 4, 5}
	target := 7
	res := twoSumHashMapSinglePass(nums, target)
	fmt.Println(res)
}