package main

import (
	"fmt"
)

func worker() {

	arr := []int{1, 6, 4, 10, 2, 5}
	n := len(arr)
	fmt.Print("_ ,")

	for i := 1; i < n; i++ {

		for j := i - 1; j >= 0; j-- {

			if arr[j] < arr[i] {

				fmt.Print(arr[j], ",")
				break
			}

		}

	}

}

func main() {

	worker()

}
