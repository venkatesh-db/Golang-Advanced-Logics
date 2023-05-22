//fibonaciee series

package main

import "fmt"

func fibo(x int, y int) chan int {
	c := make(chan int)
	go func() {
		for i, j := x, y; ; i, j = j, i+j {
			c <- i
		}
	}()
	return c
}

func main() {
	geneartor := fibo(0, 1)
	for i := 0; i < 5; i++ {
		fmt.Println(<-geneartor)
	}
}
