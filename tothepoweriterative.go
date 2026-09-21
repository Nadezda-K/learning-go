package sprint

// package main
// import "fmt"


func ToThePowerIterative(n int, power int) int  {
	if n < 0 {
		return 0
	}
	// Factorial of 0 equals 1
	if power == 0 {
		return 1
	}

    result := n 
    for power > 1 {
        result *= result
        power--
    }

    return result
}

// func main() {
// 	fmt.Println( ToThePowerIterative(2, 6) )
// }
