package main

import "fmt"

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	for i, _ := range nums { 
		product := 1 
		for j, valJ := range nums { 
			if j == i { continue }
			product *= valJ 
		} 
	  	output[i] = product 
	}

	return output
}

func main() {
	nums := []int{1,2,4,6}
	res := productExceptSelf(nums)
	fmt.Println(res)
}