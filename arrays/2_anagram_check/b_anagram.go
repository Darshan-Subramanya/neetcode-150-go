package main

import "fmt"

func anagramCheck(s string, t string) bool {
	if len(s) != len(t) { return false }

	count := [26]int{}

	for i, val := range s {
        count[val-'a']++
        count[t[i]-'a']--
    }

    for _, val := range count {
        if val != 0 {
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
