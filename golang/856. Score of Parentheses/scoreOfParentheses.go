// package scoreOfParentheses
package main

import (
	"fmt"
)

/*
Example 1:

Input: "()"
Output: 1
Example 2:

Input: "(())"
Output: 2
Example 3:

Input: "()()"
Output: 2
*/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func scoreOfParentheses(S string) int {
	stack := []int{}
	stack = append(stack, 0)
	for _, c := range S {
		if c == '(' {
			stack = append(stack, 0)
		} else {
			l := len(stack)
			v, w := stack[l-1], stack[l-2]
			stack = stack[:l-2]
			stack = append(stack, w+max(2*v, 1))
		}
	}
	return stack[len(stack)-1]
}

// 只有()才能得分，而且嵌套的()会将得分加倍,
// 题目可以简化为求所有的()的得分，其中得分多少取决于被嵌套的层数
func scoreOfParenthesesCount(S string) int {
	res := 0
	depth := uint(0)
	last := ' '
	for _, c := range S {
		if c == '(' {
			// 嵌套层数增加
			depth++
		} else {
			// 嵌套层数减少
			depth--
			// 上一个字符为'('， 表明刚好组成一对')'，可以得分
			if last == '(' {
				res += 1 << depth
			}
		}
		last = c
	}
	return res
}

func main() {
	v := "(())"
	fmt.Println(scoreOfParentheses(v))
}
