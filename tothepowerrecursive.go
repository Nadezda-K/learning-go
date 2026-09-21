package sprint

// package main
// import "fmt"


func ToThePowerIterative(n int, power int) int  {
	if n < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

    result := n * ToThePowerIterative(n, power-1)

    return result
}

// func main() {
// 	fmt.Println( ToThePowerIterative(2, 6) )
// }
