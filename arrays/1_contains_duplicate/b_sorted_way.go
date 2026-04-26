package main


import "fmt"
import "sort"

func hasDuplicateChk(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			return true
		}
	}
	return false
}

func main() {
	var nums = []int{1,2,3,1}
	fmt.Println(hasDuplicateChk(nums))
}
