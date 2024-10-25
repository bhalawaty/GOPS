package main

import "fmt"

/*
https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	var m int = 3
	nums1 := []int{1, 2, 3, 0, 0, 0}

	var n int = 3
	nums2 := []int{2, 5, 6}

	fmt.Println(merge(nums1, m, nums2, n))

}

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	nums1ArrayLastIndex := m - 1
	nums2ArrayLastIndex := n - 1
	allNums1ArrayLastIndex := m + n - 1

	for nums1ArrayLastIndex >= 0 && nums2ArrayLastIndex >= 0 {
		if nums1[nums1ArrayLastIndex] > nums2[nums2ArrayLastIndex] {
			nums1[allNums1ArrayLastIndex] = nums1[nums1ArrayLastIndex]
			nums1ArrayLastIndex--
		} else {
			nums1[allNums1ArrayLastIndex] = nums2[nums2ArrayLastIndex]
			nums2ArrayLastIndex--
		}
		allNums1ArrayLastIndex--
	}

	for nums2ArrayLastIndex >= 0 {
		nums1[allNums1ArrayLastIndex] = nums2[nums2ArrayLastIndex]
		nums2ArrayLastIndex--
		allNums1ArrayLastIndex--
	}

	return nums1
}
