package main

import "fmt"



//Anything that matches the description right here can be called as a type of Bot!
//We define the getGreeting() for both of the structs we created.
//
type Bot interface {
 	getGreeting() string
 }

 type user struct{
 	name string
 }

 //example more syntax
 type bot interface {
 	getGreeting(string,int) (string,error)
 	getBotVersion() float64
 	respondToUser(user) string
 	//Function name
 	//list of arguments
 	//list of return types
 }

 type englishBoto struct{}
 type spanishBoto struct{}

func main()  {
	eb := englishBoto{}
	sb := spanishBoto{}

	printGreeting(eb)
	printGreeting(sb)
}

//This  is a function that works on the Bot interface, If you are a type in this program with a function called getGreeting()
// and you return a string, bot type can use you.
//We can call the function directly on bot
//We define the functions that will change inside interface and then implement it and then the common functions
//use them.
func printGreeting(bot Bot) {
	fmt.Println(bot.getGreeting())
}

// There are own logical implementation of the getGreeting functions
func (eb englishBoto) getGreeting() string{
	return "Hello there"
}

func (sb spanishBoto) getGreeting() string{
	return "Hola"
}
/**
	concrete type - map,struct, int,string,englishBot
	interface type - bot

	We can create a value of the concrete types while we cannot create values of the interface types.
	We can only create values directly out of the concrete types.
 */