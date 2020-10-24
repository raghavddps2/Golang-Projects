package main

import "fmt"

func main(){

	/**
		The crux is that the slice data structure internally uses arrays and even the copy poins to that array.
		So, even when we dont use pointers, the array is only modified.
		When we create a slice, GO wll automatically create an array and a structure that records the length of the
		slice, the capacity of the slice and a refernence to the underlying array
		value types - int,float,string,bool,struct - use pointers
		reference types - slices,maps,channels,pointers,functions - use pointers
	 */
	mySlice := []string {"Hi", "There", "How", "Are", "You"}
	fmt.Println(mySlice)
	updateFirstElement(mySlice)
	fmt.Println(mySlice)
}

func updateFirstElement(s []string) {
	s[0] = "Hola"
}
