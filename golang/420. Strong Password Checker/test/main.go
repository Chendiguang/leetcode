/*
Given a string containing just the characters '(' and ')',
find the length of the longest valid (well-formed) parentheses
substring. For "(()", the longest valid parentheses substring
is "()", which has length = 2. Another example is “)(()()))”,
where the longest valid parentheses substring is “(()())”,
which has length = 6.
*/
package main

import "fmt"

func main() {
	res := longestValidParenthesesSubstring(")(()")
	fmt.Println(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestValidParenthesesSubstring(s string) int {
	index := []int{}
	stack := []rune{}
	maxLen := 0
	for i, c := range s {
		// 遇到左括号, 将其入栈
		if c == '(' {
			stack = append(stack, c)
			index = append(index, i)
		} else {
			// 遇到右括号, 分类讨论
			// 1. 当前栈顶是左括号, 则消去并计算长度
			if len(stack) != 0 && stack[len(stack)-1] == '(' {
				curLen := 0
				// top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index = index[:len(index)-1]
				if len(stack) == 0 {
					curLen = i + 1
				} else {
					curLen = i - index[len(index)-1]
				}
				maxLen = max(maxLen, curLen)
			} else {
				// 2. 右括号, 且当前栈为空时
				stack = append(stack, c)
				index = append(index, i)
			}
		}
	}
	return maxLen
}

func longestValidParentheses(s string) int {
	maxLen, last := 0, -1
	lefts := []int{}
	for i, c := range s {
		if c == '(' {
			lefts = append(lefts, i)
		} else {
			if len(lefts) == 0 {
				// no match ')'
				last = i
			} else {
				lefts = lefts[:len(lefts)-1]
				// find one match group, calculate it
				if len(lefts) == 0 {
					maxLen = max(maxLen, i-last)
				} else {
					// satck still has '('
					maxLen = max(maxLen, i-lefts[len(lefts)-1])
				}
			}
		}
	}
	return maxLen
}
