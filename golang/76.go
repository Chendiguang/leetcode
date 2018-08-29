/*
	Given a string S and a string T, find the minimum window
	in S which will contain all the characters in T in complexity O(n).
*/
package leetcode

import "math"

func minWindow(s string, t string) string {
	// The result is s[head : head+gap]
	start, end, head, gap := 0, 0, -1, math.MaxInt32
	count := len(t)
	mp := make([]int, 256)
	for i := 0; i < count; i++ {
		mp[t[i]]++
	}

	for end < len(s) {
		mp[s[end]]--
		if mp[s[end]] >= 0 {
			count--
		}
		end++
		// count is zero means has include all char of t
		for count == 0 {
			if end-start < gap {
				gap = end - start
				head = start
			}
			mp[s[start]]++
			if mp[s[start]] > 0 {
				count++
			}
			start++
		}
	}
	if head == -1 {
		return ""
	} else {
		return s[head : head+gap]
	}
}
