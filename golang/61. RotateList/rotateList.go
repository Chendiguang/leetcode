package RotateList

/*
61. Rotate List

Given a linked list, rotate the list to the right by k places,
where k is non-negative.

Example 1:
Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	// 遍历链表得到其长度
	l, p := 1, head
	for p.Next != nil {
		p = p.Next
		l++
	}
	// 将链表首尾相接，形成一个环
	p.Next = head
	// 从head出发走l-k%l步断开即可形成期望的环结构
	k = l - k%l
	p = head
	for k > 1 {
		p = p.Next
		k--
	}
	// 断开位置作为新的头节点返回，断开位置的尾巴设置为nil
	head = p.Next
	p.Next = nil
	return head
}

// double pointers
func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	p1, p2 := head, head
	// 链表长度
	l := 0
	for head != nil {
		head = head.Next
		l++
	}
	head = p1

	k %= l
	if k == 0 {
		return head
	}
	// 寻找倒数第k+1个节点
	for i := 0; i < k; i++ {
		p2 = p2.Next
	}
	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	// 画图就容易理解了
	newHead := p1.Next
	p1.Next = nil
	p2.Next = head
	return newHead
}
