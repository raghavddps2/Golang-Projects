package main

import "fmt"

func main() {

	//Creating a slice of numbers
	numbers := []int {0,1,2,3,4,5,6,7,8,9,10}

	//Iterating through the slice
	for _,j := range numbers {
		if j%2 == 0 {
			fmt.Printf("%d is Even", j)
			fmt.Println()
		} else {
			fmt.Printf("%d is Odd", j)
			fmt.Println()
		}
	}
}
func getOddOrEven(n int) string {
	if n%2 == 0 {
		return "Even"
	} else{
		return "Odd"
	}

}
