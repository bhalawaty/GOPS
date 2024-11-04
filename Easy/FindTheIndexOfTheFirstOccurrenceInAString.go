package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle)) // Output: 2

	haystack = "aaaaa"
	needle = "bba"
	fmt.Println(strStr(haystack, needle)) // Output: -1

	haystack = "a"
	needle = ""
	fmt.Println(strStr(haystack, needle)) // Output: 0

	haystack = "sadbutsad"
	needle = "sad"
	fmt.Println(strStr(haystack, needle)) // Output: 0

	haystack = "leetcode"
	needle = "leeto"
	fmt.Println(strStr(haystack, needle)) // Output: -1

}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	lastNeedleIndex := 0

	for index, _ := range haystack {
		lastNeedleIndex = index + needleLen
		if lastNeedleIndex > len(haystack) {
			return -1
		}
		result := haystack[index:lastNeedleIndex]
		if result == needle {
			return index
		}
	}
	return -1
}
