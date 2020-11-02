// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// Follow up: Could you implement a solution with a linear runtime complexity and without using extra memory?

package main

func singleNumber(nums []int) int {
	result := 0
	for i := range nums {
		result ^= nums[i]
	}
	return result
}
