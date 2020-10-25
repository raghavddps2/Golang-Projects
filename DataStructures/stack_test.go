package main

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	stack := stack{}
	stack.top = -1

	stack.push(10)

	//Now, if the stack is not empty this tells push is success
	if stack.isEmpty() == false {
		fmt.Println("Test Pass")
	} else{
		fmt.Errorf("Test for Push function failed")
	}
}

func TestPop(t *testing.T){
	stack := stack{}

	stack.push(10)
	stack.push(20)

	if stack.peek() != 20 {
		fmt.Errorf("Test failed for Pop function")
	} else{
		fmt.Println("Test Pass")
	}

}

func TestIsEmpty(t *testing.T) {
	stack := stack{}
	if stack.isEmpty() {
		fmt.Println("Test Pass")
	} else {
		fmt.Println("Test failed for IsEmpty function")
	}
}

func TestPeek(t *testing.T){
	stack := stack {}
	stack.push(20)
	if stack.peek() != 20 {
		fmt.Errorf("Test failed for Peek function")
	} else {
		fmt.Println("Test Pass")
	}
}
