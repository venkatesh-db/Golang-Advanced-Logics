//odd /even using time.sleep

package main

import (
	"fmt"
)

func oddandeven(n int, ch chan int) {

	if n%2 == 0 {

		fmt.Println(n, "the number is even ")
	} else {
		fmt.Println(n, "the number is odd ")
	}
	ch <- n
}

func main() {
	fmt.Println("even and odd")
	var ch1 chan int = make(chan int)
	go oddandeven(8, ch1)
	fmt.Println(<-ch1)

	go oddandeven(11, ch1)
	fmt.Println(<-ch1)

}
