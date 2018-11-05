package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next // tmp作为中间节点，记录当前节点的下一个节点的位置
		cur.Next = prev // 当前节点指向前一个节点
		prev = cur      // prev移到当前节点
		cur = tmp       // 指针后移，处理下一个节点
	}
	return prev
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head // 将后一个链表节点指向前一个节点
	head.Next = nil       // 将原链表中前一个节点指向后一个节点的关系断开
	return newHead
}

// 就地反转
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 设定一个头结点
	dummy := new(ListNode)
	dummy.Next = head

	// 设置两个指针
	prev := dummy.Next
	pCur := prev.Next
	for pCur != nil {
		prev.Next = pCur.Next  // prev连接下一次需要反转的节点
		pCur.Next = dummy.Next // 反转节点pCur
		dummy.Next = pCur      // 纠正头结点dummy的指向
		pCur = prev.Next       // pCur指向下一次要反转的节点
	}
	return dummy.Next
}
