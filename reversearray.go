// 40
// reverse array
package main

import "fmt"

func reverse() {
	arr := []int{12, 14, 5, 78, 90, 82}
	j := len(arr) - 1
	var temp int

	for i := 0; i < len(arr)/2; i++ {

		temp = arr[i]
		arr[i], arr[j] = arr[j], temp
		j = j - 1

	}
	fmt.Println(arr)

}

func main() {
	reverse()

}
