package main

import "fmt"

/**
	Program to find the sum of numbers from 1 to n
 */
func main(){
	n := 10
	sum := 0
	for i:= 1;i<=n;i++ {
		sum += i
	}
	fmt.Println(sum)
}
