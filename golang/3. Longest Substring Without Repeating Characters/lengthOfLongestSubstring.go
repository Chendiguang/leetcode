package lengthOfLongestSubstring

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	// 保存最近出现过的字符的索引
	usedChar := make(map[rune]int)
	// 初始化最长子字符串的起始坐标和长度
	var start, maxLen int
	for i, char := range s {
		if v, ok := usedChar[char]; ok && start <= v {
			start = v + 1
		} else {
			maxLen = max(maxLen, v-start+1)
		}
		usedChar[char] = i
	}
	return maxLen
}
