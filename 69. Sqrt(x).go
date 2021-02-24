package main

func mySqrt(x int) int {
	var i int = 1
	for ; i <= x; i++ {
		if i*i > x {
			break
		}
	}
	i -= 1
	return i
}
