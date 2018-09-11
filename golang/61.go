package leetcode

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Rotate the singly-linked list with two pointers
func rotateRight(head *ListNode, k int) *ListNode {
	if head.Next == nil || k == 0 {
		return head
	}

	start, tmp := head, head
	l := 0
	for head != nil {
		head = head.Next
		l++
	}
	head = start
	k %= l
	for k > 0 {
		head = head.Next
		k--
	}

	for head.Next != nil {
		head = head.Next
		tmp = tmp.Next
	}
	head.Next = start
	start = tmp.Next
	tmp.Next = nil
	return start
}

// AC解法
func rotateRightAc(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 遍历链表，得到链表长度l，链表首尾节点相连，形成一个环
	p := head
	l := 1
	for p.Next != nil {
		p = p.Next
		l++
	}
	p.Next = head

	// 根据环的结构，可以知道：从head开始走l - k%l步的位置处断开，
	// 形成新的链表就是想要的链表。
	k = l - k%l
	p = head
	for k > 1 {
		p = p.Next
		k--
	}

	// 首尾处理，断开位置作为新的头节点返回，断开位置的尾巴设置为nil
	// 简单点说，就是把当前点的指针指向head，
	head = p.Next
	p.Next = nil
	// p = head
	return head
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 遍历链表，得到其长度, 并对k取模
	l := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		l++
	}
	k %= l
	if k == 0 {
		return head
	}
	// 寻找倒数第k+1个节点
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p2 = p2.Next
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// 将head开头的，和现在p1指向的开头调换
	newHead := p1.Next
	p1.Next = nil
	p2.Next = head
	return newHead
}
