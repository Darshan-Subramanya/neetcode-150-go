package main

import "fmt"

func anagramCheck(s string, t string) bool {
	if len(s) != len(t) { return false }

	dict := make(map[rune]int)

	for _, val := range s {
		dict[val]++
	}

	for _, val := range t {
		dict[val]--
	}

	for _, count := range dict {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	res := anagramCheck(s, t)
	if res {
		fmt.Printf("%s and %s are anagrams\n", s, t)
	} else {
		fmt.Println("Not an anagram")
	}
}