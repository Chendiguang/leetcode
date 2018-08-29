package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

var tmp *ListNode

func isPalindrome(head *ListNode) bool {
	tmp = head
	return isPalindromeList(head)
}

func isPalindromeList(node *ListNode) bool {
	if node == nil {
		return true
	}
	res := isPalindromeList(node.Next) && (tmp.Val == node.Val)
	tmp = tmp.Next
	return res
}
