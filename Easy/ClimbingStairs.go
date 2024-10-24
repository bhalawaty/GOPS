package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	//fmt.Println(climbStairs(3))
	//fmt.Println(climbStairs(4))
	//fmt.Println(climbStairs(5))
	//fmt.Println(climbStairs(6))
	//fmt.Println(climbStairs(7))

}

/*
*using  Multiple Assignment or Tuple Assignment.
The right side of the assignment (secondStep, firstStep + secondStep)
is evaluated before any changes to the variables occur
*/
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	first := 1
	second := 1

	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}

	return second
}

/**
please see another solution and spot the deff please
*/
//func climbStairs(n int) int {
//	if n <= 1 {
//		return 1
//	}
//
//	firstStep := 0
//	secondStep := 1
//	for i := 1; i <= n; i++ {
//		temp := firstStep
//		firstStep = secondStep
//		secondStep = temp + secondStep
//	}
//	return secondStep
//}
