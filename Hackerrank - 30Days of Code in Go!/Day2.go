package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var i1 uint64
	var d1 float64
	var s1 string
	// Read and save an integer, double, and String to your variables.
	fmt.Scanln(&i1)
	fmt.Scanln(&d1)
	//To scan the text
	if scanner.Scan(){
		s1 = scanner.Text()
	}
	// Print the sum of both integer variables on a new line.
	fmt.Println(i+i1)
	// Print the sum of the double variables on a new line.
	//THis is how we print the float values onto the console.
	fmt.Printf("%.1f",d+d1)

	fmt.Println()
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s+s1)
}