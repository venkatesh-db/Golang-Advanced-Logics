//odd /even using time.sleep

package main

import (
	"fmt"
	"time"
)

func even(n int) {

	if n%2 == 0 {

		fmt.Println("the number is even ")
	} else {
		fmt.Println("the number is odd ")
	}
}

func main() {
	fmt.Println("even and odd")

	go even(5)
	time.Sleep(1)

}
