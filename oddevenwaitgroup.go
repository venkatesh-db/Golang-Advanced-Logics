package main

import (
	"fmt"
	"sync"
)

func print_odd(odd chan int) {
	for {
		select {
		case n := <-odd:
			fmt.Println("Odd: ", n)
		}
	}
}

func print_even(even chan int) {
	for {
		select {
		case n := <-even:
			fmt.Println("Even: ", n)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	odd := make(chan int)
	even := make(chan int)

	go print_odd(odd)
	go print_even(even)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
	wg.Done()
	wg.Done()
	wg.Wait()
}
