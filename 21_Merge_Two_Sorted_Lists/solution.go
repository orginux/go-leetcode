package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	a := list1
	b := list2

	head := ListNode{}
	curr := &head

	for a != nil && b != nil {
		if a.Val < b.Val {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}
		curr = curr.Next
	}
	if a != nil {
		curr.Next = a
	}
	if b != nil {
		curr.Next = b
	}
	return head.Next

}
