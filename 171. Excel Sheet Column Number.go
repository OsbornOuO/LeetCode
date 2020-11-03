package main

import (
	"fmt"
	"math"
)

func titleToNumber(s string) int {
	total := 0
	for i := range s {
		if len(s)-1 == i {
			total += int(s[i] - 64)
		} else {
			fmt.Println()
			total += int(math.Pow(26, float64(len(s)-i-1))) * int(s[i]-64)
		}
	}
	return total
}
