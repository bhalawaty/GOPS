package main

import "fmt"

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
	k := 0
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
			count = 1
		} else {
			count++
			if count <= 2 {
				k++
				nums[k] = nums[i]
			}
		}
	}
	return k + 1 // k+1 is the number of unique elements because indices start at 0
}
