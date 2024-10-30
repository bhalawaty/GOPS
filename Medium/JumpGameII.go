package main

import "fmt"

/*
https://leetcode.com/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2,3,0,1,4}

	fmt.Println(canJump(nums))

}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] == 0 {
		return 0
	}

	lIndex := len(nums) - 1
	maxJump := 0
	Jumps := 0
	currJump := 0
	for i, val := range nums {

		maxJump = max(val+i, maxJump)

		if maxJump >= lIndex {
			return Jumps + 1
		}

		if i == currJump { //all the trick here we need to just track the maxJump each time and save it as a current jump
			Jumps++
			currJump = maxJump
		}

	}
	return Jumps
}
