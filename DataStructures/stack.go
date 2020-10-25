package main

import "fmt"

func main() {
	stack := stack{}
	stack.top = -1
	//We push the values

	fmt.Println(stack.isEmpty())
	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.push(40)
	stack.push(50)
	stack.push(60)

	fmt.Println(stack.peek())

	fmt.Println(stack.pop())
	fmt.Println(stack.peek())

	fmt.Println(stack.isEmpty())
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()

	fmt.Println(stack.pop())
	fmt.Println(stack.peek())
	fmt.Println(stack.isEmpty())
}
