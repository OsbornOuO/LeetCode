// Given a non-empty array of integers, return the k most frequent elements.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
    dict := make(map[int]int)
    for _, v := range nums {
        dict[v]++
    }

    var keys []int
    
    for k, _ := range dict {
        keys = append(keys, k)
    }
    
    sort.Slice(keys, func (i int, j int) bool {
        return dict[keys[i]] > dict[keys[j]]
    })
    
    return keys[:k]
}
Ｆ