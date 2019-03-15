// package coinChange
package main

import "fmt"

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/

func coinChange(coins []int, amount int) int {
	// dp[i]表示凑齐钱数i所需的最少硬币个数
	// 状态转移方程: dp[i] = min(dp[i], dp[i - coins[j]] + 1)
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i >= coin && dp[i-coin] != -1 && (dp[i] == -1 || dp[i-coin]+1 < dp[i]) {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	coins := []int{2}
	amount := 1
	res := coinChange(coins, amount)
	fmt.Println(res)
}
