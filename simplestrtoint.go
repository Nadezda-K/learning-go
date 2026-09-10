//package sprint
package main
import "fmt"

//func SimpleStrToInt(s string) int {
func main() {
    s := "0000000010203"
    for i:=0; i<len(s); i++ {
        fmt.Println(s[i], string(s[i]))
    }
}
