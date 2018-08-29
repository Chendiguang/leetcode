package leetcode

/*
844. Backspace String Compare
Given two strings S and T,
return if they are equal when both are typed into empty text editors.
# means a backspace character.
*/

// 模拟栈的用法， O(M+N)
func backspaceCompare(S string, T string) bool {
	return build(S) == build(T)
}

func build(s string) string {
	res := []rune{}
	for _, c := range s {
		if c != '#' {
			res = append(res, c)
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}
