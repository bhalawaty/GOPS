package main

import "fmt"

/*
*
link => https://leetcode.com/problems/roman-to-integer/?envType=study-plan-v2&envId=top-interview-150
test cases :
"III"
"LVIII"
"MCMXCIV"
*/
func main() {

	// Example 1
	s := "III"
	fmt.Println(romanToInt(s))

	// Example 2
	s = "LVIII"
	fmt.Println(romanToInt(s))

	// Example 3
	s = "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var romanNumerals = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	length := len(s)
	for i := 0; i < length; i++ {
		value := romanNumerals[s[i]]

		if i+1 < length && romanNumerals[s[i+1]] > value {
			total -= value
		} else {
			total += value
		}
	}
	return total
}
