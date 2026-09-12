package main

import (
        "fmt"
//        "bufio"
//        "os"
//        "strconv"
//        "string"
        )

func main() {
    fmt.Println("Welcome to the Cypher Tool!")

    // User input. Define encrypt or decrypt operation
    var operation int
    fmt.Println("Select operation (1/2):")
    fmt.Println("1. Encrypt.")
    fmt.Println("2. Decrypt.")
    
    fmt.Scan(&operation)

    // User input. Choose encription algorithm
    var cypher int
    fmt.Println("Select cypher (1/3):")
    fmt.Println("1. ROT13.")
    fmt.Println("2. Reverse.")
    fmt.Println("3. Optional")
    
    fmt.Scan(&cypher)








    fmt.Printf("You choose operation %02d and cypher %02d \n", operation, cypher)
}
