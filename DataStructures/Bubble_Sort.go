package main

import "fmt"

func main(){
	arr := []int{5,4,3,2,1}
	fmt.Printf("UnSorted array is %v ",arr)
	bubbleSort(arr)
	fmt.Println()
	fmt.Printf("Sorted Array is %v ",arr)
}

//When we pass particular values, we need to pass it as pointers and not really just direct values.
func swap(value1 *int, value2 *int){
	temp := *value1
	*value1 = *value2
	*value2 = temp
}
/**
	5 4 3 2 1
	i = 0
		j = 0 to j = 3
			arr[0] > arr[1] 4 5 3 2 1
			arr[1] > arr[2] 4 3 5 2 1
			arr[2] > arr[3] 4 3 2 5 1
			arr[3] > arr[4] 4 3 2 1 5
	i = 1
		j = 0 to j = 2
			arr[0] > arr[1] 3 4 2 1 5
			arr[1] > arr[2] 3 2 4 1 5
			arr[2] > arr[3] 3 2 1 4 5
 */
func bubbleSort(arr []int) {

	// Looping over from 0 to len(arr) - 1
	for i:=0;i<len(arr)-1;i++ {

		//On every iteration, we will move the largest element to the last.
		for j:=0;j<len(arr)-i-1;j++ {
			if arr[j] > arr[j+1] {
				//We swap the values if arr[j] > arr[j+1]
				//We pass the addresses and not the values.
				swap(&arr[j],&arr[j+1])
			}
		}
	}

}