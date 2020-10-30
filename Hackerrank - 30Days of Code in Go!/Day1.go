package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	//To take an entire line as string input, we can use the bufio.NewReader and it basically takes os.Stdin as the reader.
	reader := bufio.NewReader(os.Stdin)

	//We read the string sepearated by the delimiter of a new line
	text,_ := reader.ReadString('\n')

	//We first print the hello world and then we go ahead and print the text we just read.
	fmt.Println("Hello, World.")
	fmt.Println(text)

}