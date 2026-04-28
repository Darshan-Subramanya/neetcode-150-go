package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7
	res := twoSum(nums, target)
	fmt.Println(res)
}