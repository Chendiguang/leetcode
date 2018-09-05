package Partition_Labels

/*
A string S of lowercase letters is given.
We want to partition this string into as many parts as possible
so that each letter appears in at most one part,
and return a list of integers representing the size of these parts.
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
*/

func partitionLabels(S string) []int {
	// 用一个数组保存每个元素出现的最右边的位置
	last := [26]int{}
	for i, c := range S {
		last[c-'a'] = i
	}
	j, anchor := 0, 0
	ans := []int{}
	for i, c := range S {
		if j < last[c] {
			j = last[c]
		}
		// 此时说明：对这个划分已经到了结尾
		if i == j {
			ans = append(ans, j-anchor+1)
			anchor = j + 1
		}
	}
	return ans
}
