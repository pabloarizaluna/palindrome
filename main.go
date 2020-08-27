package main

import (
	"fmt"
	"os"
)

func isPalindrome(str string) bool {
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		if str[start] != str[end] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This command must have a word to check.\n./palindrome [word]")
		return
	}
	fmt.Println(isPalindrome(os.Args[1]))
}
