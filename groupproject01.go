
package main

import (
        "fmt"
//        "bufio"
//        "os"
//        "strconv"
//        "string"
        )

func input_operation() bool {
    // User input. Define encrypt or decrypt operation
    var op int
    
    fmt.Println("Select operation (1/2):")
    fmt.Println("1. Encrypt.")
    fmt.Println("2. Decrypt.")
    
    fmt.Scan(&op)
    if ( op != 1 ) && ( op != 2 ) {
        fmt.Println("\n Incorrect input. Please choose 1 or 2.\n")
        input_operation()
    }

    var bool_op bool
    if ( op == 1 ) {
        bool_op = true
    }
    if ( op == 2 ) {
        bool_op = false
    }
    return bool_op
}

func input_cypher() string {
    // User input. Choose encription algorithm
    var c int
    fmt.Println("Select cypher (1/2/3):")
    fmt.Println("1. ROT13.")
    fmt.Println("2. Reverse.")
    fmt.Println("3. Optional")
    
    fmt.Scan(&c)
    if ( c != 1 ) && ( c != 2 ) && ( c != 3 ){
        fmt.Println("\n Incorrect input. Please choose 1, or 2, or 3 .\n")
        input_cypher()
    }

    var str_c string
    switch c {
        case 1 :
            str_c = "ROT13"
        case 2 :
            str_c = "Reverse"
        case 3 :
            str_c = "Optional"
    }
    return str_c
}




func main() {
    fmt.Println("Welcome to the Cypher Tool!\n")

    // User input. Define encrypt or decrypt operation
    is_encryption := input_operation()
    fmt.Println("is_encryption", is_encryption)

    // User input. Choose encription algorithm
    cypher := input_cypher()
    fmt.Println("cypher algorithm", cypher)

    // User input. Meassage
    var message string
    fmt.Println("Enter the message:")
    
    fmt.Scan(&message)





//    switch operation {
//        case 1 :
//             str_operation := "Encrypted"
//        case 2 :
//             str_operation := "Dencrypted"
//    }

//    fmt.Printf("You choose operation %02d and cypher %02d \n", operation, cypher)
//    fmt.Printf("Decrypted message using reverse%v:\n", returned_message)

}
