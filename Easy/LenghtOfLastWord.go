package main

import "fmt"

/*
*
link => https://leetcode.com/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150
test cases
"Hello World"
"   fly me   to   the moon  "
"luffy is still joyboy"
"a"
"a "
*/
func main() {

	// Example 1
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))

	// Example 2
	s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))

	// Example 3
	s = "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 1 {
		return 1
	}

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if s[i] == ' ' && count > 0 {
			return count
		} else {
			count = 0
		}
	}
	return count
}
