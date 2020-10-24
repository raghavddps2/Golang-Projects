package main

import "fmt"

/**
	This is a unction to find whether a string is a pallindrome or not
 */
func main() {
	str := "xyyyyx"
	ans := findPallindrome(str)
	fmt.Println(ans)
}

func findPallindrome(str string) bool {
	for i := 0 ; i < len(str)/2 ;i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}