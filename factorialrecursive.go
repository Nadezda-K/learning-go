package sprint

// package main
// import "fmt"


func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}
	// Factorial of 0 equals 1
	if n == 0 {
		return 1
	}

	// Find biggest possible int
	maxInt := int(^uint(0) >> 1)

    result := n * FactorialRecursive(n-1) // solution by recursive function

    if result > maxInt/n {
        return 0
    } 

    return result
}

// func main() {
// 	fmt.Println( FactorialRecursive(21) )
// }
