package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main(){
	response,error := http.Get("https://example.com")
	if error != nil {
		fmt.Println("Error :", error)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw,response.Body)
}


//Now, logWriter implements the Writer Interface
func (logWriter) Write(bs [] byte) (int,error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes")
	return len(bs),nil
}