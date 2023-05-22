package main

import "fmt"

func IsBalanced2(text string) bool {
	isBalanced := true
	bMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	s := make([]rune, 0, len(text))

	for _, c := range text {

		if _, ok := bMap[c]; ok == true {
			s = append(s, c)
		} else if len(s) == 0 {
			isBalanced = false
			break
		} else if bMap[s[len(s)-1]] != c {
			isBalanced = false
			break
		} else {
			s = s[:len(s)-1]
		}
	}

	if len(s) != 0 {
		isBalanced = false
	}

	return isBalanced
}

func main() {
	fmt.Println(IsBalanced2("[{()}]"))
	fmt.Println(IsBalanced2("[{()]"))
}
