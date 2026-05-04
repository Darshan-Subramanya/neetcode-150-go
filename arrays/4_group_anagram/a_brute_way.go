package main

import "fmt"

func isAnagram(a, b string) bool {
    dict := make(map[rune]int)
    for _, val := range a {
        dict[val]++
    }
    for _, val := range b {
        dict[val]--
    }
    for _, count := range dict {
        if count != 0 {
            return false
        }
    }
    return true
}

func groupAnagrams(strs []string) [][]string {
	visited := make(map[int]bool)
    result := [][]string{}

    for i := 0; i < len(strs); i++ {
		if !visited[i] {
			group := []string{strs[i]}
			visited[i] = true
			for j := i + 1; j < len(strs); j++ {
				if !visited[j] && isAnagram(strs[i], strs[j]) {
					group = append(group, strs[j])
					visited[j] = true
				}
			}
			result = append(result, group)
		}
    }

    return result
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	res := groupAnagrams(strs)

	fmt.Println(res)
}
