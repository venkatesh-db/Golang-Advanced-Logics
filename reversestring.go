// reverse string
package main

import "fmt"

func main() {
	s1 := "Hemanth Kumar"
	l1 := len(s1)
	s2 := ""
	i := 0
	for ; i < l1; i++ {
		tempword := " "
		for ; i < l1 && string(s1[i]) != " "; i++ {
			tempword = string(s1[i]) + tempword
		}
		s2 = s2 + tempword + ""
	}
	fmt.Print(s2)
}
