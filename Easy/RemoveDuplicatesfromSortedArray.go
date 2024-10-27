package main

import "fmt"

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{3, 2, 2, 3}

	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {

	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			nums[k+1] = nums[i]
			k++
		}
	}
	return k + 1 // k+1 is the number of unique elements because indices start at 0

}
