package main

import "fmt"

type stack struct {
	top int
	stackVal [1000]int
}

func (stack *stack) push(value int) bool {
	if stack.top >= 999 {
		fmt.Println("Stack Overflow")
		return false
	} else {
		stack.top += 1
		stack.stackVal[stack.top] = value
		return true
	}
}

func (stack *stack) pop() int {
	if stack.top == - 1{
		fmt.Println("Stack Is Empty")
		return -1
	} else{
		ans := stack.stackVal[stack.top]
		stack.top -= 1
		return ans
	}
}

func (stack *stack) isEmpty() bool {
	if stack.top == - 1 {
		return true
	}	else {
		return false
	}
}

func (stack *stack) peek() int {
	if !stack.isEmpty() {
		return stack.stackVal[stack.top]
	}
	return -1
}