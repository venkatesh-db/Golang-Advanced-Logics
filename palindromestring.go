//palindrome using string

package main

import (
	"fmt"
	"strings"
)

func main() {

	var os string = "Madam"
	var rs string = ""
	var n = len(os)

	for i := n - 1; i >= 0; i-- {
		rs += string(os[i])
	}

	if strings.ToUpper(os) == strings.ToUpper(rs) { //converts all characters to upper case or lower case
		fmt.Println("The given string is Palindrome")
	} else {
		fmt.Println("The given string is NOT a Palindrome")
	}
}
