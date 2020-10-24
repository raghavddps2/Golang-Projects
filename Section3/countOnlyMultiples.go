package main
import "fmt"

/**

	This is a program that finds the sum of all numbers divisble by 3 and 5 between 1 to 10
	and the sum of numbers not divisible by both
 */
func main(){
	n := 10
	sum1 := 0
	sum2 := 0
	for i:= 1; i<=n ;i++ {
		if i%3 == 0 || i%5 == 0 {
			sum1 += i
		} else{
			sum2 += i
		}
	}
	fmt.Println("Sum of all numbers divisible by 3 and 5 from 1 to 10 are: ", sum1)
	fmt.Println("Sum of all numbers not divisible by 3 and 5 from 1 to 10 are ",sum2)
}
