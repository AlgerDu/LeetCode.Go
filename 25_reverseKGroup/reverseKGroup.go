package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head
	}

	root := &ListNode{0, head}
	head = root
	offset := 0

	for head.Next != nil {
		h, t, enough := reverseK(head.Next, k-1, &offset)

		if enough {
			head.Next = h

			head = t
			offset = 0
		} else {
			break
		}
	}

	return root.Next
}

func reverseK(node *ListNode, k int, offset *int) (*ListNode, *ListNode, bool) {

	*offset++

	if node.Next == nil && *offset <= k {
		return nil, nil, false
	}

	if k == *offset {
		head := node.Next
		tail := node

		tail.Next = head.Next
		head.Next = tail

		return head, tail, true
	} else {
		head, tail, enough := reverseK(node.Next, k, offset)

		if enough {
			tmp := tail.Next
			tail.Next = node
			node.Next = tmp

			return head, node, enough
		} else {
			return node, tail, enough
		}
	}
}
