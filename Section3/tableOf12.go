package main
import "fmt"

/**
	This is used to print a table of 12
 */
func main(){
	n := 12
	for i:=1 ;i <= 10; i++ {
		fmt.Println( n," X " ,i," =", i*n)
	}
}

