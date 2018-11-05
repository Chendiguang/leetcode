package restoreIpAddresses

import (
	"strconv"
)

/*
Given a string containing only digits,
restore it by returning all possible valid IP address combinations.

Example:

Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
*/

// DFS
func restoreIpAddresses(s string) []string {
	var ips []string
	if len(s) < 4 || len(s) > 12 {
		return ips
	}
	// 必须传递指针
	helper(s, 0, 0, &ips)
	return ips
}

func helper(s string, offset, dotNum int, ips *[]string) {
	l := len(s)
	if offset >= l {
		return
	}
	if dotNum == 3 {
		// 每一位都不能有0x, 0xx这种格式
		if l-offset >= 2 && s[offset] == '0' {
			return
		}
		if i, err := strconv.Atoi(s[offset:l]); i <= 255 && err == nil {
			*ips = append(*ips, s)
		}
		return
	}
	if s[offset] == '0' {
		t := s[:offset+1] + "." + s[offset+1:]
		helper(t, offset+2, dotNum+1, ips)
		return
	}
	for i := 1; i <= 3; i++ {
		if offset+i > l {
			return
		}
		if v, err := strconv.Atoi(s[offset : offset+i]); v <= 255 && err == nil {
			t := s[:offset+i] + "." + s[offset+i:]
			helper(t, offset+i+1, dotNum+1, ips)
		}
	}
}
