package removeNthFromEnd

/*
Given a linked list, remove the n-th node from the end of list
and return its head.

Example:
Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end,
the linked list becomes 1->2->3->5.

one pass

常用技巧：
1. Dummy head：简化改变、删除头指针的处理。
2. 前后双指针：多用于链表反转。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// two pass, 4ms
// 先用一个runner指针走n步，然后再用一个walker指针从head开始和此时的runner指针
// 一起向后移动，当runner移动到链表尾部时，walker指针指向的即为倒数第n个节点
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	i := 0
	walker, runner := head, head
	for runner != nil && i < n {
		runner = runner.Next
		i++
	}
	if i < n {
		return nil
	}
	if runner == nil {
		return head.Next
	}

	for runner.Next != nil {
		walker = walker.Next
		runner = runner.Next
	}
	walker.Next = walker.Next.Next
	return head
}

// 0ms, Sample as removeNthFromEnd2 but more faster
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	// 利用一个虚拟节点，减少边界条件的判断
	dummy := ListNode{0, head}
	l, r := &dummy, &dummy

	// r 先走n步
	for i := 0; r != nil && i < n+1; i++ {
		r = r.Next
	}
	// l, r 一起往后走，直到r走到链表末尾
	for r != nil {
		l, r = l.Next, r.Next
	}
	l.Next = l.Next.Next
	head = dummy.Next
	return head
}
