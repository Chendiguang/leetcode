/*
A password is considered strong if below conditions are all met:

It has at least 6 characters and at most 20 characters.
It must contain at least one lowercase letter, at least one uppercase letter, and at least one digit.
It must NOT contain three repeating characters in a row ("...aaa..." is weak, but "...aa...a..." is strong, assuming other conditions are met).
Write a function strongPasswordChecker(s), that takes a string s as input, and return the MINIMUM change required to make s a strong password. If s is already strong, return 0.

Insertion, deletion or replace of any one character are all considered as one change.
*/

package strongPasswordChecker

import (
	"math"
)

func max(res ...int) int {
	m := math.MinInt32
	for _, n := range res {
		if n > m {
			m = n
		}
	}
	return m
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func strongPasswordChecker(s string) int {
	// counting missing character
	missing := [3]int{}
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			missing[0]++
		} else if c >= 'A' && c <= 'Z' {
			missing[1]++
		} else if c >= '0' && c <= '9' {
			missing[2]++
		}
	}
	miss := 0
	for i := 0; i < len(missing); i++ {
		if missing[i] == 0 {
			miss++
		}
	}
	// counting repeats
	p, n := 2, len(s)
	one, two, repeat := 0, 0, 0
	for p < n {
		if s[p] == s[p-1] && s[p] == s[p-2] {
			x := 2
			for p < n && s[p] == s[p-1] {
				x++
				p++
			}
			repeat += x / 3
			if x%3 == 0 {
				one++
			} else if x%3 == 1 {
				two++
			}
		} else {
			p++
		}
	}

	if n < 6 {
		return max(miss, 6-n)
	} else if n <= 20 {
		return max(miss, repeat)
	} else {
		deletes := n - 20
		repeat -= min(deletes, one)
		repeat -= min(max(deletes-one, 0), two*2) / 2
		repeat -= max(deletes-one-2*two, 0) / 3
		return deletes + max(miss, repeat)
	}
}
