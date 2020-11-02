// Write a program that outputs the string representation of numbers from 1 to n.

// But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

package main

import "strconv"

func fizzBuzz(n int) []string {
	tmp := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			tmp[i-1] = "FizzBuzz"
		case i%3 == 0:
			tmp[i-1] = "Fizz"
		case i%5 == 0:
			tmp[i-1] = "Buzz"
		default:
			tmp[i-1] = strconv.Itoa(i)
		}
	}
	return tmp
}
