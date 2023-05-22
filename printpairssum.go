package main

import "fmt"

func printpairs(arr []int, n int, sum int) {

	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {

			if arr[i]+arr[j] == sum {

				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}

				fmt.Print("[", arr[i], ",", arr[j], "]")

			}
		}
	}
}

func main() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6, 9}
	n := len(arr)
	sum := 10
	printpairs(arr, n, sum)
}
