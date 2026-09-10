package sprint

//import "fmt"

func Countdown(n int) string {
    str := ""
    for i:=n; i>0; i-=2 {
        str += i + ", "
    }
    str += "0!"
}
