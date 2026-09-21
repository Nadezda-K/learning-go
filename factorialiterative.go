//package sprint

package main
import "fmt"


func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}
	// Factorial of 0 equals 1
	if n == 0 {
		return 1
	}

	// Find biggest possible int
	maxInt := int(^uint(0) >> 1)

//    result := n * FactorialIterative(n-1) // solution by recursive function

    result := 1
    for n > 1 {
        result *= n
        n--
        // overflow handle
        if result > maxInt/n {
            return -1
        } 
    }

    return result
}

func main() {
	fmt.Println( FactorialIterative(5) )
}
