package main

import "fmt"

/*
https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	var val int = 3
	nums := []int{3, 2, 2, 3}

	fmt.Println(removeElement(nums, val))

}

func removeElement(nums []int, val int) int {
	k := 0
	for i, value := range nums {
		if value != val {
			if k != i { // Only assign if the current index i is different from k
				nums[k] = value
			}
			k += 1
		}
	}
	fmt.Println(nums)
	return k
}
