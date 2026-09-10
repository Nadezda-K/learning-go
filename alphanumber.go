package sprint
//package main

//import "fmt"

func AlphaNumber(n int) string {
//func main() {
    n:= -1280
    number := n
    str := ""
    var digit int

    if ( n <= 0 ) {
        number = 0 - number
        fmt.Println(number)
    }
    
    for number > 0 {
        digit = number%10
        number = number/10
        str = string(digit+97) + str
    }

    if ( n <= 0 ) {
        str = string(45) + str // rune(45) is a minus sign
    }

    return str
//    fmt.Println(str)
}
