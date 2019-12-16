package gostack

import "fmt"

type node struct {
	data interface{}
	next *node
}

type lStack struct {
	Top *node
	len int //Denotes current number of elements in the stack
	size int //Total number of elements the stack can accommodate
}

/*CreateStack creates an empty stack which can accommodate n nodes.
  Initially the length len of stack will be zero and Top pointer will point to nil.
*/
func CreateStack(n int) *lStack {
	return &lStack{
		Top:  nil,
		len:  0,
		size: n,
	}
}

func(stk *lStack) Push(data interface{}) {
	if !stk.IsFull() {
		newNode := &node{
			data: data,
			next: nil,
		}
		newNode.next = stk.Top
		stk.Top = newNode
		stk.len++
		return
	}
	fmt.Println("StackOverflow!")
}

func(stk *lStack) Pop() {
	if !stk.IsEmpty() {
		temp := stk.Top.next
		stk.Top.next = nil
		stk.Top = temp
		stk.len--
		return
	}
	fmt.Println("StackUnderFlow!")
}

func (stk *lStack) GetTop() interface{} {
	if !stk.IsEmpty() {
		return stk.Top.data
	}
	return nil
}

func (stk *lStack) Peek(pos int) (interface{}, bool) {
	temp := stk.Top
	for i := 0; temp != nil && i < pos - 1; i++ {
		temp = temp.next
	}
	if temp != nil {
		return temp.data, true
	}
	return nil, false
}

func (stk *lStack) IsFull() bool {
	//Limiting the number of elements in stack by size
	if stk.size == stk.len {
		return true
	}
	return false
}

func (stk *lStack) IsEmpty() bool {
	if stk.Top == nil {
		return true
	}
	return false
}

func Display(top *node) {
	for top != nil {
		fmt.Println(top.data)
		top = top.next
	}
}