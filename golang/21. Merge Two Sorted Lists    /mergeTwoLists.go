package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// Similar with mergeSort
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp ListNode
	res := &tmp

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Next = l1
			l1 = l1.Next
		} else {
			res.Next = l2
			l2 = l2.Next
		}
		res = res.Next
	}

	if l1 != nil {
		res.Next = l1
	} else {
		res.Next = l2
	}

	return tmp.Next
}
