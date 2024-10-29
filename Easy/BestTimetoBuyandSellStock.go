package main

import "fmt"

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

	nums := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(nums))

}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit

}
