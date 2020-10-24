package main

import "fmt"

/**
	This is a funtion t o implement binarySerach in go
 */
func main() {
	arr := []int {34,45,56,76,89,100,233}
	key1 := 233
	key2 := 34
	fmt.Println(binarySearch(arr,key1))
	fmt.Println(binarySearch(arr,key2))
}

func binarySearch(arr []int, key int) (bool,int) {
	low := 0
	high := len(arr)-1

	for low <= high {
		mid := (low+high)/2
		if arr[mid] == key {
			return true,mid
		} else if arr[mid] > key {
			high = mid - 1
		} else{
			low = mid + 1
		}
	}
	return false,-1
}
