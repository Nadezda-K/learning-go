
package main

import (
        "fmt"
        "bufio"
        "os"
//        "strconv"
        "strings"
        )

func input_operation() bool {
    // User input. Define encrypt or decrypt operation
    var op string
    
    fmt.Println("Select operation (1/2):")
    fmt.Println("1. Encrypt.")
    fmt.Println("2. Decrypt.")
    
    //fmt.Scan(&op)
    reader := bufio.NewReader(os.Stdin)
    op, _ = reader.ReadString('\n')
    op = strings.TrimSpace(op)
    if ( op != "1" ) && ( op != "2" ) {
        fmt.Println("\n Incorrect input. Please choose 1 or 2.\n")
        input_operation()
    }

    var bool_op bool
    if ( op == "1" ) {
        bool_op = true
    }
    if ( op == "2" ) {
        bool_op = false
    }
    return bool_op
}


func input_cypher() string {
    // User input. Choose encription algorithm
    var c string
    fmt.Println("Select cypher (1/2/3):")
    fmt.Println("1. ROT13.")
    fmt.Println("2. Reverse.")
    fmt.Println("3. Optional")
    
    //fmt.Scan(&c)
    reader := bufio.NewReader(os.Stdin)
    c, _ = reader.ReadString('\n')
    c = strings.TrimSpace(c)
    if ( c != "1" ) && ( c != "2" ) && ( c != "3" ){
        fmt.Println("\n Incorrect input. Please choose 1, or 2, or 3 .\n")
        input_cypher()
    }

    var str_c string
    switch c {
        case "1" :
            str_c = "rot13"
        case "2" :
            str_c = "Reverse"
        case "3" :
            str_c = "Optional"
    }
    return str_c
}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {
    // User input. Define encrypt or decrypt operation
    toEncrypt = input_operation()
    fmt.Println("is_encryption", toEncrypt)

    // User input. Choose encription algorithm
    encoding = input_cypher()
    fmt.Println("cypher algorithm", encoding)

    // User input. Meassage
    fmt.Println("Enter the message:")
    fmt.Scan(&message)

    return toEncrypt, encoding, message
}

func is_letter(r rune) bool {
    // Checking if rune is a letter
    if ( r > rune('a') && r < rune('z') ) || ( r > rune('A') && r < rune('Z') ) {
        return true
    }
    return false
}

func ShiftBy(r rune, step int) string {
    return string((int(r) - int('a') + step) % 26 + int('a'))
}



// Encrypt the message with rot13
func encrypt_rot13(s string) string {
    str := ""
    for i, v := range s {
        str += ShiftBy(rune(v), 13)
        fmt.Println(i, v, )
    }
    return str
}

// Encrypt the message with reverse
//func encrypt_reverse(s string) string {}

// Decrypt the message with rot13
//func decrypt_rot13(s string) string {}

// Decrypt the message with reverse
//func decrypt_reverse(s string) string {}




func main() {
    fmt.Println("Welcome to the Cypher Tool!\n")

    // Get the input data required for the operation
    is_encryption, cypher, message := getInput()
    is_encryption, cypher, message = true, "rot13", "Hello, World! 1234"

    fmt.Printf("In main : is_encryption =  %v\n cypher %v\n message %v \n", is_encryption, cypher, message)
    
    var str_result string

    switch cypher {
    case "rot13" :
        if is_encryption {
            str_result = encrypt_rot13(message)
        }
//        if is_encryption {
//            decrypt_rot13(message)
//        }
//    case "Reverse" :
//        if is_encryption {
//            encrypt_reverse(message)
//        }
//        if is_encryption {
//            decrypt_reverse(message)
//        }
    //case "Optional":

    }

    fmt.Println(str_result)

}
