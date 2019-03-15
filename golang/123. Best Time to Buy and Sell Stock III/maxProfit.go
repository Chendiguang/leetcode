// package maxProfit

package main

import "fmt"

/*
Design an algorithm to find the maximum profit.
You may complete at most two transactions.
Input: [3,3,5,0,0,3,1,4]
Output: 6

Input: [1,2,3,4,5]
Output: 4

Input: [7,6,4,3,1]
Output: 0
*/

func maxProfit1(prices []int) int {
	if prices == nil || len(prices) <= 2 {
		return 0
	}
	k, n := 2, len(prices)
	local, global := make([]int, k+1), make([]int, k+1)
	for i := 0; i < n-1; i++ {
		diff := prices[i+1] - prices[i]
		for j := k; j > 0; j-- {
			local[j] = max(global[j-1], local[j]) + diff
			global[j] = max(local[j], global[j])
		}
	}
	return global[2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp[k, i] = max(dp[k, i-1], prices[i] - prices[j] + dp[k-1, j-1]), j=[0..i-1]
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	dp := [][]int{}
	k := 3
	for i := 0; i < k; i++ {
		dp = append(dp, make([]int, len(prices)))
	}

	fmt.Println(dp)
	for i := 1; i < k; i++ {
		minPrice := prices[0]
		for j := 1; j < len(prices); j++ {
			// fmt.Printf("i: %d, j: %d\n", i, j)
			minPrice = min(minPrice, prices[j]-dp[i-1][j-1])
			dp[i][j] = max(dp[i][j-1], prices[j]-minPrice)
		}
	}
	fmt.Println(dp)
	return dp[k-1][len(prices)-1]
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}
