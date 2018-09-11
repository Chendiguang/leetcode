package partitionList

type ListNode struct {
	Val  int
	Next *ListNode
}

// 用两个链表分别连接小于和大于等于x 的结点，
// 然后再把两个结点链接到一起
func partition(head *ListNode, x int) *ListNode {
	// 用h1保存小于的，h2保存大于的
	// dummy1指向h1的头部，dummy2指向h2的头部
	h1, h2 := &ListNode{}, &ListNode{}
	dummy1, dummy2 := h1, h2

	for head != nil {
		if head.Val < x {
			h1.Next = head
			h1 = h1.Next
		} else {
			h2.Next = head
			h2 = h2.Next
		}
		head = head.Next
	}

	h2.Next = nil
	h1.Next = dummy2.Next
	return dummy1.Next
}
