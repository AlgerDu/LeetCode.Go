package leetcode

import "testing"

func TestExample(t *testing.T) {

	ll1 := []int{2, 4, 3}
	ll2 := []int{5, 6, 4}

	l1 := &ListNode{ll1[0], nil}
	l2 := &ListNode{ll2[0], nil}

	offset1 := l1
	offset2 := l2

	for i := 1; i < len(ll1); i++ {
		offset1.Next = &ListNode{ll1[i], nil}
		offset1 = offset1.Next
	}

	for i := 1; i < len(ll2); i++ {
		offset2.Next = &ListNode{ll2[i], nil}
		offset2 = offset2.Next
	}

	rst := addTwoNumbers(l1, l2)

	t.Logf("%d", rst.Val)
}

func TestExample2(t *testing.T) {

	ll1 := []int{0}
	ll2 := []int{7, 3}

	l1 := &ListNode{ll1[0], nil}
	l2 := &ListNode{ll2[0], nil}

	offset1 := l1
	offset2 := l2

	for i := 1; i < len(ll1); i++ {
		offset1.Next = &ListNode{ll1[i], nil}
		offset1 = offset1.Next
	}

	for i := 1; i < len(ll2); i++ {
		offset2.Next = &ListNode{ll2[i], nil}
		offset2 = offset2.Next
	}

	rst := addTwoNumbers(l1, l2)

	t.Logf("%d", rst.Val)
}
