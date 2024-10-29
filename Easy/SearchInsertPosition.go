package main

import "fmt"

/*
https://leetcode.com/problems/search-insert-position/description/
*/
func main() {

	nums := []int{1, 3, 5, 6}
	target := 5

	//nums := []int{1, 3, 5, 6}
	//target := 2
	//
	//nums := []int{1, 3, 5, 6}
	//target := 7

	fmt.Println(searchInsert(nums, target))

}

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := 0
	for low <= high {
		mid = (high + low) / 2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return low

}
