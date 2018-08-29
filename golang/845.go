package leetcode

/*
845. Longest Mountain in Array
Let's call any (contiguous) subarray B (of A) a mountain if the following properties hold:

B.length >= 3
There exists some 0 < i < B.length - 1 such that B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
(Note that B could be any subarray of A, including the entire array A.)

Given an array A of integers, return the length of the longest mountain.
Return 0 if there is no mountain.
*/
// Can you solve it using only one pass?
// Can you solve it in O(1) space?

func longestMountain(A []int) int {
	if A == nil || len(A) < 3 {
		return 0
	}
	base, ans, n := 0, 0, len(A)
	for base < n {
		end := base
		if end+1 < n && A[end] < A[end+1] {
			// 从当前点出发，直到遇到峰值
			for end+1 < n && A[end] < A[end+1] {
				end++
			}
			// 峰值
			if end+1 < n && A[end] > A[end+1] {
				// for循环从峰值向右边扩张，直到遇到边界点
				for end+1 < n && A[end] > A[end+1] {
					end++
				}
				ans = Max(ans, end-base+1)
			}
		}
		// 存在山脉，则下一个base点为当前的end，即base=end
		// 否则，base=base+1
		base = Max(end, base+1)
	}
	return ans
}
