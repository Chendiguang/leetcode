package minWindow

/*
Given a string S and a string T, find the minimum window in S
which will contain all the characters in T in complexity O(n).

Example:
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"

Note:
If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

// list
func minWindow(s string, t string) string {
	lt := len(t)
	// 遍历t, 把对应的字符及其出现的次数保存到数组
	mp := make([]int, 256)
	for i := 0; i < lt; i++ {
		mp[t[i]]++
	}
	// head mean the start index of result, start+gap is the end.
	start, end, head, gap := 0, 0, -1, len(s)+1
	for end < len(s) {
		// 在遍历过程中，对映射mp的元素进行判断
		mp[s[end]]--
		if mp[s[end]] >= 0 {
			lt--
		}
		end++
		// lt=0时，符合t长度要求了，开始更新左边界start
		for lt == 0 {
			if end-start < gap {
				gap = end - start
				head = start
			}
			mp[s[start]]++
			if mp[s[start]] > 0 {
				lt++
			}
			start++
		}
	}
	if head == -1 {
		return ""
	}
	return s[head : gap+head]
}
