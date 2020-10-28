package main

import (
	"fmt"
	"net/http"
	"time"
)
/**
	This program is currently running sequentially, it makes a request, waits for response to come back and
	then it moves forward.

	Maybe we can think of a kind of a parallel approach to this problem, with that we can make every request right away
	rather than waiting for the previous request to finish and then moving forward.
 */
func main() {
	websites := []string {
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//This is how we create a brand new channel.
	c := make(chan string)

	for i:= 0;i<len(websites);i++ {
		//we send the channel in the function.
		go checkStatus(websites[i],c)
	}
	//channel <- 5, send the value into the channel.
	//myNumber <- channel. Wait for a value to be sent to the channel.
	//This is a blocking line of code, waits for the result to be printed.
	//Receving message from a channel is a blocking thing!

	//We spawned up 5 routines and we want 5 results.
	//This puts up a loop expecting the 5 messages from the routines through the channels
	//for i:=0;i<len(websites);i++ {
	//	//Till it receives a message from the channel, it blocks the for loop from going further.
	//	//This will keep on checking the website one by one by one and so on..
	//	go checkStatus(<-c,c)
	//	i = i - 1
	//}
	//Other infinite loop way will be
	//for {
	//	go checkStatus(<-c,c)
	//}

	//We take every element out of the channel and then call the checkStatus on that link.
	//This goes on till we keep on receiving the channel message
	//We should never trybto share variables and just use channels.
	//for l:= range c {
	//	//This pauses the current goRoutine, here main GoRoutine
	//	go checkStatus(l,c)
	//}

	for l:= range c {
		//This is called as a function literal - Anonymous or lambda function.
		//A function literal is an unnamed function
		//Passing l as an argument and accepting it.
		go func(l string) {
			time.Sleep(time.Second*5)
			go checkStatus(l,c)
		}(l)
	}
}
/**
	Routines interact with each other through Channels.
 */
func checkStatus(website string,c chan string){
	//time.Sleep(time.Second*5)
	_,error := http.Get(website)
	if error != nil {
		fmt.Println(website + " is down")
		c <- website
	} else {
		fmt.Println(website + " is up")
		c <- website
	}
}
/**
	Go routine : It is a separate line of code execution that can be used to handle non blocking code.
	Purpose of channel - This is used for communication among go routines
	I did not do function literals properly!
 */