package sprint

//package main
//import "fmt"

func SimpleStrToInt(s string) int {
    result := int(0)
    beginning_is_zero := true
    for i:=0; i<len(s); i++ {
        if ( rune(s[i]) < rune('0') ) || ( rune(s[i]) > rune('9') ) {
            return 0
        }
        if ( rune(s[i]) > rune('0') ) && ( beginning_is_zero == true ) {
            beginning_is_zero = false
        }
        if ( beginning_is_zero == false ) {
            result = result*10 + int(s[i]-'0')
        }
    }
    return result
}

/*
func main() {
    x := SimpleStrToInt("102039")

    fmt.Printf("The function returned: %d\n", x)
}
*/
