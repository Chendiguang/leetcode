package maxProfit

/*
可以无限次买入卖出
Input: [7,1,5,3,6,4]
Output: 7
Explanation:
Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
*/

// 贪心法。从前向后遍历数组，只要当天的价格高于前一天的价格，就算入收益
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			maxprofit += tmp
		}
	}
	return maxprofit
}
