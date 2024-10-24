package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.com/problems/longest-common-prefix/description/
*/
func main() {
	arr := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(arr)
	fmt.Println(result)
}

/*
*
the syntax [:i] is used to slice a sequence up to, but not including, the index i.
*/
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
