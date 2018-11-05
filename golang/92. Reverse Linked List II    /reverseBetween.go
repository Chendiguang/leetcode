package reverseBetween

/*
Reverse a linked list from position m to n.
Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	// 指针走到m
	for i := 1; i < m; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	pre := head
	// 记录m, n指向的点
	var front, last *ListNode
	last = pre.Next
	for i := m; i <= n; i++ {
		head = pre.Next
		pre.Next = head.Next
		head.Next = front
		front = head
	}
	head = pre.Next
	pre.Next = front
	last.Next = head
	return dummy.Next
}

// traverse list to the mth node, then use two pointer p1 and p2 to reverse the nodes in [m,n].
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy

	// 先连接之前的m-1个元素
	for i := 0; i < m-1; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.Next
	}

	// // 用两个指针协助[m, n]范围的翻转
	// p1, p2 := cur.Next, cur.Next.Next
	// for ; i < n-1; i++ {
	// 	p1.Next = p2.Next
	// 	p2.Next = cur.Next
	// 	cur.Next = p2
	// 	p2 = p1.Next
	// }
	p := cur.Next
	for i := 0; i < n-m; i++ {
		tmp := p.Next
		p.Next = tmp.Next
		tmp.Next = cur.Next
		cur.Next = p
	}
	return dummy.Next
}
