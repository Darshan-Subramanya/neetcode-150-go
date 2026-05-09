package main

import (
	"fmt"
	"strconv"
)

func Encode(strs []string) string {
	res := ""
    for _, val := range strs {
        res += fmt.Sprintf("%d#%s", len(val), val)
    }
    return res
}

func Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		l, _ := strconv.Atoi(encoded[i:j])
		word := encoded[j+1 : j+1+l]
		res = append(res, word)
		i = j + 1 + l
	}
	return res
}

func main() {
	dummy_input := []string{"Hello", "World"}
	encode := Encode(dummy_input)
	fmt.Println(encode)
	decode := Decode(encode)
	fmt.Printf("%v\n", decode)
}
