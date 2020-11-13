// Given a set of distinct integers, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

// Example:

// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := [][]int{[]int{}}

	for _, v := range nums {

		for _, val := range ans {
			fmt.Println(val, v)
			temp := append([]int{}, val...)
			temp = append(temp, v)
			ans = append(ans, temp)
		}
		fmt.Println()
	}

	return ans
}
