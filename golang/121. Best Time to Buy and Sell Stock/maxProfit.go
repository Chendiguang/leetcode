package maxProfit

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction
(i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.
Input: [7,1,5,3,6,4]
Output: 5
Input: [7,6,4,3,1]
Output: 0
*/

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	curMinPrice := prices[0]
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		// curMinPrice记录遍历过的最低价格
		if prices[i] < curMinPrice {
			curMinPrice = prices[i]
			// 计算当前值和最小买入价格的差值，然后比较更新
		} else if prices[i]-curMinPrice > maxprofit {
			maxprofit = prices[i] - curMinPrice
		}
	}
	return maxprofit
}

// 局部最优和全局最优， 速度慢一点
func maxProfit1(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	local, global := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		local = max(local+prices[i+1]-prices[i], 0)
		global = max(global, local)
	}
	return global
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
