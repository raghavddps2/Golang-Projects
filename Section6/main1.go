package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main(){
	eb := englishBot{}
	//sb := spanishBot{}

	printGreetings(eb)
	//printGreetings(sb)
}

func (eb englishBot) getGreetings() string{
	//This is the custom logic
	return "Hello there!!"
}

func (sb spanishBot) getGreetings() string {
	//This is the custom logic
	return "Holaa"
}

// We need to condense the two functions into one. Interfaces will solve the problem.

func printGreetings(eb englishBot) {
	fmt.Println(eb.getGreetings())
}

//func printGreetings(sb spanishBot) {
//	fmt.Println(sb.getGreetings())
//}
