package main

/**

	Maps in Go

	1. All the keys must be of same type
	2. All the values in the map must be of the same type

	With structures, we need to define all the properties or basically analogous keys beforehand
	while we dont necessarily need to do this in maps.

	to delete a key from map - > delete(colors,"red")

	Difference between Maps and Stucts

	1. In maps all keys must be of the same type while in structs all the values can be of different types.
	2. This is reference type, whereas structs are map type
	3. This is used to represent a collection of related properties, whereas in structs we need to know values at the
		compile time

 */
import "fmt"
func main() {
	/**
		This tells a map where both keys and values are strings
		Method1 to declare a map
	 */
	colors := map[string]string {
		"red" : "#ff0000",
		"green":"#4bf745",
	}
	fmt.Println(colors)

	/**
		Method2 to declare a map
	 */
	var colors1 map[string]string
	/**
		Method3 to declare a map
	 */
	colors2 := make(map[string]string)
	colors2["green"] = "#er3344"
	fmt.Println(colors1)
	fmt.Println(colors2)

	delete(colors,"red")
	fmt.Println(colors)

	actress := make(map[string]string)

	actress["raghav"] = "priyanka"
	actress["harshit"] = "deepika"

	fmt.Println(actress)
	printMap(actress)

	array := []int {1,1,1,1,2,2,3,2,1,2,4,5,6,6,4}
	countOfEachValue(array)
}
/**
	We need to define the key type as string and the value type string as well.
 */
func printMap(c map[string]string) {
	for key,value := range c {
		fmt.Println(key,value)
	}
}

func countOfEachValue(array []int) {
	resultMap := make(map[int]int)
	for i:=0 ; i< len(array); i ++ {
		value,ok := resultMap[array[i]]
		if ok {
			resultMap[array[i]] = value + 1
		} else {
			resultMap[array[i]] = 1
		}
	}

	for key,value := range resultMap {
		fmt.Println("count of ",key," is ",value)
	}
}
