package main

import "fmt"

func firstUniqChar(s string) int {
	type tmpStruct struct {
		Index int
		Count int
	}
	tmp := make(map[byte]*tmpStruct)

	for i := range s {
		fmt.Printf("%s %d\n", string(s[i]), s[i])

		if _, ok := tmp[s[i]]; !ok {
			tmp[s[i]] = &tmpStruct{
				Index: i,
				Count: 1,
			}
		} else {
			tmp[s[i]].Count++
		}
	}

	var index int
	check := false
	for i := range tmp {
		fmt.Println(i, tmp[i])
		if (!check || index >= tmp[i].Index) && tmp[i].Count == 1 {
			check = true
			index = tmp[i].Index
		}
	}

	if check {
		return index
	}
	return -1
}
