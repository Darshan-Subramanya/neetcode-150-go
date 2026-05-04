package main

import (
	"fmt"
	"sort"
)

func groupAnagramsFixedArr(strs []string) [][]string {
	res := make(map[string][]string)
	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)
		res[key] = append(res[key], word)
	}
	result := [][]string{}
	for _, val := range res {
		result = append(result, val)
	}
	return result
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	res := groupAnagramsFixedArr(strs)

	fmt.Println(res)
}
