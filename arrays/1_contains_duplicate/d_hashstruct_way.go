package main

import "fmt"

func hasDuplicateHashSetStruct(nums []int) bool {
	seen := make(map[int]struct{})
	for _, val := range nums {
		seen[val] = struct{}{}
	}
	return len(seen) < len(nums)
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(hasDuplicateHashSetStruct(nums))
}