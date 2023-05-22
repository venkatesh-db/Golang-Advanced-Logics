

package main
import "fmt"

func isDivBy2PowerM(n uint, m uint) bool {

	// if expression results to 0, then
	// n is divisible by pow(2, m)

       //        1248 163264128 
        // n :  1010 0000
	// m :  0100 0000
	// 1 :  1000 0000 

         // n :  1010 0000
	 //      0010 0000
            
	if (n & ((1 << m) - 1)) == 0 {
		return true
	}
	// n is not divisible
	return false
}

func main() {

	var n uint = 5
	var m uint = 2

	if isDivBy2PowerM(n, m) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
