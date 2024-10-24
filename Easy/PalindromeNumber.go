package main

import "fmt"

/*
Time Complexity O(N)
*/
func main() {
	result := IsPalindromeNumber(234)
	fmt.Println(result)
}

func IsPalindromeNumber(x int) bool {

	if x < 0 {
		return false
	}

	original := x
	reversed := 0
	for x != 0 {
		remainder := x % 10
		reversed = reversed*10 + remainder
		x /= 10
	}
	return original == reversed

}
