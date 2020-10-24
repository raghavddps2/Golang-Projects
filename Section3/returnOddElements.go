package main

import "fmt"

func main() {
	arr := []int {1,2,3,4,5,6,7,8,9,10}
	evenAns := findEvenNumbers(arr)
	oddAns := findOddNumbers(arr)

	fmt.Println(evenAns)
	fmt.Println(oddAns)
}

func findEvenNumbers(arr []int) []int {
	ans := []int{}
	for _,j := range arr {
		if j%2 == 0 {
			ans = append(ans,j)
		}
	}
	return ans
}

func findOddNumbers(arr []int) []int {
	ans := []int{}
	for _,j := range arr {
		if j%2 == 1 {
			ans = append(ans,j)
		}
	}
	return ans
}