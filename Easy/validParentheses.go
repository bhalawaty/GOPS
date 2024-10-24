package main

import "fmt"

func main() {
	input := "()"
	result := isValid(input)
	fmt.Println(result)

	input1 := "()[]{}"
	result1 := isValid(input1)
	fmt.Println(result1)

	input2 := "(]"
	result2 := isValid(input2)
	fmt.Println(result2)

	input3 := "([])"
	result3 := isValid(input3)
	fmt.Println(result3)
}

func isValid(s string) bool {
	var stack []rune

	parentheses := map[rune]rune{
		')': '(',
		'}': '{',
		']': '['}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != parentheses[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}
