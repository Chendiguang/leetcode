package swapPairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		tmp := pre.Next.Next
		pre.Next.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
		pre = tmp.Next
	}
	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	t := head.Next
	head.Next = swapPairs2(head.Next.Next)
	t.Next = head
	return t
}
