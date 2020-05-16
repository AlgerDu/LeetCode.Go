package solution

import (
	"strconv"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	listStr := "1,2,3,4,5"
	desireStr := "2,1,4,3,5"
	k := 2

	list := CreateList(listStr)
	desire := CreateList(desireStr)

	output := reverseKGroup(list, k)

	h1 := desire
	h2 := output
	i := 0

	for h2 != nil {
		if h2.Val != h1.Val {
			t.Errorf("%d val is %d, not %d", i, h2.Val, h1.Val)
		}

		h1 = h1.Next
		h2 = h2.Next
		i++
	}
}

func TestExample2(t *testing.T) {
	listStr := "1,2,3,4,5"
	desireStr := "3,2,1,4,5"
	k := 3

	list := CreateList(listStr)
	desire := CreateList(desireStr)

	output := reverseKGroup(list, k)

	h1 := desire
	h2 := output
	i := 0

	for h2 != nil {
		if h2.Val != h1.Val {
			t.Errorf("%d val is %d, not %d", i, h2.Val, h1.Val)
		}

		h1 = h1.Next
		h2 = h2.Next
		i++
	}
}

func TestExample3(t *testing.T) {
	listStr := "1,2,3,4,5"
	desireStr := "1,2,3,4,5"
	k := 1

	list := CreateList(listStr)
	desire := CreateList(desireStr)

	output := reverseKGroup(list, k)

	h1 := desire
	h2 := output
	i := 0

	for h2 != nil {
		if h2.Val != h1.Val {
			t.Errorf("%d val is %d, not %d", i, h2.Val, h1.Val)
		}

		h1 = h1.Next
		h2 = h2.Next
		i++
	}
}

func CreateList(listStr string) *ListNode {

	head := &ListNode{0, nil}
	node := head

	nodeStrs := strings.Split(listStr, ",")

	for _, v := range nodeStrs {
		nodeVal, _ := strconv.Atoi(v)
		node.Next = &ListNode{nodeVal, nil}

		node = node.Next
	}

	return head.Next
}
