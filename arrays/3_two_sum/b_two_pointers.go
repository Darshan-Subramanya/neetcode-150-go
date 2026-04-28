package main

import (
	"sort"
	"fmt"
)

func twoSumtwoPointers(nums []int, target int) []int {
	pairsArr := make([][2]int, len(nums))

	// adding to pairsArr
	for i, val := range nums {
		pairsArr[i] = [2]int{val, i}
	}

	// sorting array
	sort.Slice(pairsArr, func(i, j int) bool {
		return pairsArr[i][0] < pairsArr[j][0]
	})

	// initialize the pointers
	leftPtr := 0
	rightPtr := len(nums) - 1

	for leftPtr < rightPtr {
		cur := pairsArr[leftPtr][0] + pairsArr[rightPtr][0]
		if cur == target {
			i, j := pairsArr[leftPtr][1], pairsArr[rightPtr][1]
			if i < j {
				return []int{i, j}
			}
			return []int{j, i}
		} else if cur > target {
			rightPtr--
    	} else {
        	leftPtr++
    	}
	}

	return nil
}


func main() {
	nums := []int{6, 7, 3, 4, 5}
	target := 7
	res := twoSumtwoPointers(nums, target)
	fmt.Println(res)
}
