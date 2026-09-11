package sprint

//package main
//import "fmt"

func GCD(a, b int) int {
    var gcd int
    
    if ( a > b) {
        a,b = b,a // swap values
    }
    
    if ( a == 0 ) && ( b == 0 ) { 
        gcd = b 
    }
    
    for i:=a; i>0; i-- {
        if ( b%i == 0 ) && ( a%i == 0 ) {
            gcd = i
            break
        }
    }

    return gcd
}

/*
func main() {
    fmt.Println("Greatest Common Divisor = ", GCD(11, 7))
}
*/

