package main
import "fmt"

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	maximum := 0
	for _,j := range arr {
		if j > maximum {
			maximum = j
		}
	}
	fmt.Println("The maximum element in the array is : ",maximum)

}
