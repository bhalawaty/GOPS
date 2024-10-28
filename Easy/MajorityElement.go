package main

import "fmt"

/*
https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{2, 2, 1, 1, 1, 2, 2, 1, 1}

	fmt.Println(majorityElement(nums))

}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exists := m[nums[i]]; exists {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}

	i := 0
	result := 0
	for index, val := range m {
		if val >= i {
			i = val
			result = index
		}
	}
	return result
}
