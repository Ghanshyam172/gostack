package gostack

import "fmt"

type stack struct {
	size int
	top int
	slc []interface{}
}

//Creates a stack of given size with top initialized to -1
func CreateStackUsingArray(size int) *stack{
	return &stack{
		size: size,
		top:  -1,
		slc:  make([]interface{}, size),
	}
}

func (stk *stack) PushArrStack(data interface{}) {
	if !stk.IsFull() {
		stk.top++
		stk.slc[stk.top] = data
		return
	}
	fmt.Println("Stack Overflow! Stack is full")
}

//Removes top element from stack
//element will be there in the slice, only top will be decremented, element will be replaced by other element when a new
//element comes
func (stk *stack) PopArrStack() {
	if !stk.IsEmpty() {
		stk.top--
		return
	}
	fmt.Println("Stack Underflow! No element to delete")
}

//return top element of the stack
func (stk *stack) TopArrStack() interface{} {
	return stk.slc[stk.top]
}

//return given element from the stack, starting from 0 to size-1 , if not found returns nil, false
func (stk *stack) Peek(pos int) (interface{}, bool) {
	if pos > stk.size || stk.top == -1 {
		return nil, false
	}
	return stk.slc[stk.top - pos + 1], true
}

func (stk *stack) IsEmpty() bool {
	if stk.top == -1 {
		return true
	}
	return false
}

func (stk *stack) IsFull() bool {
	if stk.top == stk.size - 1 {
		return true
	}
	return false
}


func DisplayArrStack(stk *stack) {
	for i := stk.top; i >= 0; i-- {
		fmt.Println(stk.slc[i])
	}
}

