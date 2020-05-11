package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	tmp := 0

	offset1 := l1
	offset2 := l2

	for offset1.Next != nil || offset2 != nil {

		if offset1 != nil {
			tmp += offset1.Val
		}

		if offset2 != nil {
			tmp += offset2.Val
		}

		offset1.Val = tmp % 10
		tmp = tmp / 10

		if offset1.Next == nil {
			offset1.Next = &ListNode{0, nil}
		}

		offset1 = offset1.Next

		if offset2 != nil {
			offset2 = offset2.Next
		}
	}

	if offset1 != nil {
		tmp += offset1.Val
	}

	for tmp > 0 {
		offset1.Next = &ListNode{0, nil}
		offset1 = offset1.Next
		offset1.Val = tmp % 10

		tmp = tmp / 10
	}

	return l1
}
