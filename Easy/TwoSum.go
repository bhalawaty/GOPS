package main

import "fmt"

///*
//Time Complexity O(N)
//*/
//func twoSum(nums []int, target int) []int {
//	numMap := make(map[int]int)
//
//	for i, num := range nums {
//		numMap[num] = i
//	}
//
//	for indexNum, numValue := range nums {
//		complement := target - numValue
//		if indexNumMap, exists := numMap[complement]; exists && indexNum != indexNumMap {
//			return []int{indexNum, indexNumMap}
//		}
//	}
//
//	return nil // Return nil if no solution is found
//}

/*
Time Complexity O(N)
*/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for indexNum, numValue := range nums {
		complement := target - numValue
		if indexNumMap, exists := numMap[complement]; exists {
			return []int{indexNum, indexNumMap}
		} else {
			numMap[numValue] = indexNum
		}
	}

	return nil // Return nil if no solution is found
}

/*
*
link => https://leetcode.com/problems/two-sum/description/
*/
func main() {
	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1)) // [0, 1]

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2)) // [1, 2]

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3)) // [0, 1]
}
