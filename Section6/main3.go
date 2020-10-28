package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
/**
	interfaces are not generic types - Other languages have 'generic' types, go doesn't have.
	interfaces are implicit - We don't manually have to say that our custom type satisfies interfaces
	interfaces are contact to help us manage types
	interfaces are tough. Step #1 is how to understand them.

	HTTP request to google.com -> Print response to the terminal

 */

func main() {
	fmt.Println("Hello World")
	httpCall()
}

func httpCall() {
	//The GET Method basically returns the response and error, here error will be nill
	response,error := http.Get("https://mraghav.tech")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println(response)
	//Number we want to initialize the slice with
	//bs := make([]byte,99999)
	//response.Body.Read(bs)
	//fmt.Println(string(bs))
	//
	io.Copy(os.Stdout,response.Body)

}

/**
	Response struct{
		status string
		statusCode int
		Body io.ReadCloser
	}

	//Now, we definitely need to talk about the io.ReadCloser interface
	type ReadCloser interface{
		Reader
		Closer
	}

	type Reader interface{
		Read([]byte) (int,error)
	}

	type Closer interface {
		Close() (error)
	}

 */
