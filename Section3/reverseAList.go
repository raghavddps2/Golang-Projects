package main
import "fmt"

/**
	This si a simple program to reverse an array.
 */
func main() {
	arr := []int {1,2,3,4,5,6}
	for i:=0 ;i< len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp
	}

	fmt.Println(arr)
}
