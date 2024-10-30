package main

import "fmt"

/*
https://leetcode.com/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(canJump(nums))

}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}

	lIndex := len(nums) - 1
	maxJ := 0
	for i, _ := range nums {
		if i > maxJ {
			return false
		}

		if maxJ >= lIndex {
			return true
		}
		if nums[i]+i > maxJ {
			maxJ = nums[i] + i
		}
	}
	return false
}
