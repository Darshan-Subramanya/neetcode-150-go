package main

import "fmt"

func topKFrequentBucketSort(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	result := []int{}
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, num := range bucket[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func main() {
	nums := []int{1,2,2,3,3,3}
	k := 2

	res := topKFrequentBucketSort(nums, k)
	fmt.Println(res)
}