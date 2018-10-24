package addTwoNumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		// add first, then carry 1 when over 9
		res := &ListNode{l1.Val + l2.Val, addTwoNumbers(l1.Next, l2.Next)}
		if res.Val > 9 {
			res = &ListNode{res.Val - 10, addTwoNumbers(res.Next, &ListNode{1, nil})}
		}
		return res
	}
}
