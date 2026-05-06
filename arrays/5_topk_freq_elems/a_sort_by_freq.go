package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {       
		freqMap[num]++  
	}

	keys := make([]int, 0, len(freqMap)) 
	for k := range freqMap { 
		keys = append(keys, k) 
	} 
	sort.Slice(keys, func(i, j int) bool { 
		return freqMap[keys[i]] > freqMap[keys[j]]
	})
	return keys[:k]
}

func main() {
	nums := []int{1,2,2,3,3,3}
	k := 2

	res := topKFrequent(nums, k)
	fmt.Println(res)
}