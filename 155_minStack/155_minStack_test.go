package leetcode

import (
	"strconv"
	"testing"
)

func Execute(t *testing.T, cmds, values []string) {
	miniStack := Constructor()

	for i := 0; i < len(cmds); i++ {
		switch cmds[i] {
		case "MinStack":
			{
				break
			}

		case "push":
			{
				v, _ := strconv.Atoi(values[i])
				miniStack.Push(v)
				break
			}

		case "pop":
			{
				miniStack.Pop()
				break
			}

		case "top":
			{
				v, _ := strconv.Atoi(values[i])
				top := miniStack.Top()
				if v != top {
					t.Errorf("stack top %d not equals %d", top, v)
				}

				break
			}

		case "getMin":
			{
				v, _ := strconv.Atoi(values[i])
				top := miniStack.GetMin()
				if v != top {
					t.Errorf("stack min %d not equals %d", top, v)
				}

				break
			}
		}
	}
}

func TestExample(t *testing.T) {

	cmds := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}
	values := []string{"", "-2", "0", "-3", "-3", "0", "", "-2"}

	Execute(t, cmds, values)
}

func TestExample_5(t *testing.T) {
	cmds := []string{"MinStack", "push", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin", "pop", "getMin"}
	values := []string{"", "2", "0", "3", "0", "0", "", "0", "", "0", "", "2"}

	Execute(t, cmds, values)
}
