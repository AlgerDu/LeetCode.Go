package leetcode

type MinStack struct {
	Head *stackNode
}

type stackNode struct {
	Val  int
	Next *stackNode

	Small *stackNode
	Big   *stackNode
}

func Constructor() MinStack {
	head := &stackNode{0, nil, nil, nil}

	stack := MinStack{head}

	return stack
}

func (this *MinStack) Push(x int) {
	node := &stackNode{x, nil, nil, nil}

	node.Next = this.Head.Next
	this.Head.Next = node

	offset := this.Head

	for offset.Big != nil && offset.Big.Val < x {
		offset = offset.Big
	}

	if offset.Big == nil {
		offset.Big = node
		node.Small = offset
	} else {
		offset.Big.Small = node
		node.Big = offset.Big

		offset.Big = node
		node.Small = offset
	}
}

func (this *MinStack) Pop() {
	node := this.Head.Next

	this.Head.Next = node.Next

	if node.Big == nil {
		node.Small.Big = nil
	} else {
		node.Big.Small = node.Small
		node.Small.Big = node.Big
	}
}

func (this *MinStack) Top() int {
	return this.Head.Next.Val
}

func (this *MinStack) GetMin() int {
	return this.Head.Big.Val
}
