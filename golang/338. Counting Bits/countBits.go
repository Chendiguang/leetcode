package countBits

/*
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:

Input: 2
Output: [0,1,1]
Example 2:

Input: 5
Output: [0,1,1,2,1,2]
*/
// 循环的在每个数 后面加0, 1即可得到接下来的数
// 故此处的状态转移方程为: res[i] = res[i/2] + i%2
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}
