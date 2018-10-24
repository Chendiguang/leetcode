package longestPalindrome

import (
	"bytes"
	"fmt"
)

func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}

	var optLeft, optRight, maxLen = 0, 0, 0

	for i := 0; i < len(s); i++ {
		// 事实上，left可以去掉，因为在整个过程left都不变，现在只是为了逻辑清晰
		left, right := i, i
		for right < len(s)-1 && s[right+1] == s[left] {
			right++
		}
		// j表示从left或right向两边扩展的长度
		j := 1
		for left-j >= 0 && right+j < len(s) && s[left-j] == s[right+j] {
			j++
		}
		curLen := right - left + 1 + 2*(j-1)
		// fmt.Printf("i = %d, j = %d, left = %d, right = %d, curLen=%d\n",
		// 	i, j, left, right, curLen)
		for maxLen < curLen {
			optLeft = left - j + 1
			optRight = right + j - 1
			maxLen = curLen
		}
		// 避免重复
		i = right
	}
	return s[optLeft : optRight+1]
}

// 马拉车算法，Time Complexity:O(n), Space Complexity: O(n)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	bs := process(s)
	var maxCenterIndex, center, right int
	p := make([]int, len(bs))
	// 保存以当前字符为中心的最大回文串的半径
	for i := 1; i < len(bs)-1; i++ {
		// j为i关于中心点center的对称点
		j := 2*center - i
		// equals to P[i] = (R > i) ? min(R-i, P[j) : 0;
		if right > i {
			if right-i > p[j] {
				p[i] = p[j]
			} else {
				p[i] = right - i
			}
		}
		// Attempt to expand palindrome centered at i
		for bs[i+1+p[i]] == bs[i-1-p[i]] {
			p[i]++
		}
		// If palindrome centered at i expand past right,
		// adjust center based on expanded palidrome.
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}

		// Find the maximum element in p
		if p[maxCenterIndex] < p[i] {
			maxCenterIndex = i
		}
	}
	start := (maxCenterIndex - 1 - p[maxCenterIndex]) / 2
	end := start + p[maxCenterIndex]
	p = p[:0]
	return s[start:end]
}

func process(s string) []byte {
	var buf bytes.Buffer
	buf.WriteString("^")
	for _, b := range s {
		buf.WriteRune('#')
		buf.WriteRune(b)
	}
	buf.WriteRune('#')
	buf.WriteRune('$')
	fmt.Println(buf.String())
	return buf.Bytes()
}
