package main

import "fmt"

/*
https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println(rotate(nums, k))

}

/*
*
Why k = k % n?
This operation is essential to handle cases where the rotation count k is greater than the length of the array n.
Rotating an array of length n by n steps results in the array looking exactly as it started.
This means a rotation of n steps is equivalent to no rotation at all.
So, when k is greater than n, we only need to rotate k % n times,
which gives the remainder of k divided by n.

Example:
Consider nums = [1, 2, 3, 4, 5, 6, 7] with n = 7 (the length of nums).
If k = 10, then rotating 10 times is effectively the same as rotating 10 % 7 = 3 times
because every 7 rotations, the array returns to its original configuration.
*/
func rotate(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	temp := make([]int, n)

	copy(temp, nums)
	newIndex := 0
	for i := 0; i < n; i++ {
		if (i + k) >= n { // we can make it like [(i + k) % n] will get same result by i will keep it like this for simplicity
			newIndex = (i + k) - n
		} else {
			newIndex = i + k
		}
		nums[newIndex] = temp[i]
	}
	return nums
}
