package middleNode

type ListNode struct {
	Val  int
	Next *ListNode
}

// one pass, fast and slow pointer
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// two pass
func middleNode2(head *ListNode) *ListNode {
	var length int
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur != nil {
		cur = cur.Next
		length++
	}
	cur = dummy
	m := (length + 1) / 2
	for i := 0; i < m; i++ {
		cur = cur.Next
	}
	return cur
}
