package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(arr)
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
		return strs[0]
	}

	sort.Strings(strs)
	fmt.Println(strs)

	length := len(strs)
	first := strs[0]
	last := strs[length-1]
	for i := range strs[0] {
		if first[i] != last[i] {
			return first[:i]
		}
	}
	return strs[0]
}
