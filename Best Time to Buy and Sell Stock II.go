// Say you have an array prices for which the ith element is the price of a given stock on day i.

// Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

// Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

package main

func maxProfit(prices []int) int {
	var tmp int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			tmp += prices[i] - prices[i-1]
		}
	}
	return tmp
}
