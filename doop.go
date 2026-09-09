package sprint

func Doop(a int, op string, b int) int {
    switch op {
        case "+":
            return a + b
        case "-":
            return a - b
        case "/" :
            if (b == 0) {
              return 0
            }
            return a / b
        case "*":
            return a * b
        case "%" :
            if ( b == 0 ): {
              return 0
            }
            return a % b

        default:
            return 0
    }
}

/*
package main

import "fmt"

func main() {
    a := int(5)
    op := string("+")
    b := int(3)
    switch op {
        case "+":
            fmt.Println(a + b)
        case "-":
            fmt.Println(a - b)
        case "/":
            fmt.Println(a / b)
        case "*":
            fmt.Println(a * b)
        case "%":
            fmt.Println(a % b)
        default:
            fmt.Println(0)
    }
}
*/
