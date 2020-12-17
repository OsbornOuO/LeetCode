package main

import (
	"sort"
)

func wiggleSort(nums []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var sNums, bNums []int
	if len(nums)%2 == 0 {
		sNums, bNums = make([]int, len(nums)/2), make([]int, len(nums)/2)
	} else {
		sNums, bNums = make([]int, len(nums)/2+1), make([]int, len(nums)/2)
	}
	copy(bNums, nums[:len(nums)/2])
	copy(sNums, nums[len(nums)/2:])

	var j, k int
	for i := range nums {
		if i%2 == 0 {
			nums[i] = sNums[j]
			j++
		} else {
			nums[i] = bNums[k]
			k++
		}
	}

	return
}
