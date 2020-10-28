package main

import (
	"fmt"
	"io"
	"os"
)

//One of the task was to take the name of the file from the command line argument.
//All we have to do is to use the os.Args utility to take the command line argument.
func main() {
	//taking the command line arguments
	args := os.Args
	fileName := args[1]
	fmt.Println(fileName)

	//opening the file
	file,error := os.Open(fileName)
	if error != nil {
		fmt.Println("Some error occured")
		os.Exit(1)
	}
	io.Copy(os.Stdout,file)
}