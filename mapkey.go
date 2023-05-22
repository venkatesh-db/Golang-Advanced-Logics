// Existing key present in map
package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "adidas", "year": "1998", "location": "bangalore"}

	_, ok1 := a["brand"] //Only checking for existing key and not its value
	_, ok2 := a["dateofbirth"]
	val1, ok3 := a["year"] //Checking for existing key and its value
	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(val1, ok3)

}
