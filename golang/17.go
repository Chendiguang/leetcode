package leetcode

// 首先，建立一个对应的数组numsMap, numsMap[i]代表数字i上面的字母
// 动态规划求解
var numMaps = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	dp(digits, 0, []byte{}, &res)
	return res
}

func dp(digits string, index int, combination []byte, result *[]string) {
	if index == len(digits) {
		*result = append(*result, string(combination))
		return
	}
	word := numMaps[digits[index]-'0']
	for i := 0; i < len(word); i++ {
		dp(digits, index+1, append(combination, word[i]), result)
	}
}
