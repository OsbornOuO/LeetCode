package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int = make([][]int, 0)
	if nums == nil || len(nums) < 4 {
		return ans
	}

	sort.Ints(nums)
	for i := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1; j < len(nums)-2; j++ {
				if j == i+1 || nums[j-1] != nums[j] {
					begin, end := j+1, len(nums)-1
					newT := target - nums[i] - nums[j]
					for begin < end {
						if nums[begin]+nums[end] == newT {
							ans = append(ans, []int{nums[i], nums[j], nums[begin], nums[end]})
							for begin < end && nums[begin] == nums[begin+1] {
								begin++
							}
							for begin < end && nums[end] == nums[end-1] {
								end--
							}
							begin++
							end--
						} else if nums[begin]+nums[end] > newT {
							end--
						} else {
							begin++
						}
					}
				}
			}
		}
	}

	return ans
}
