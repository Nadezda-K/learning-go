package sprint

//import "fmt"

func Countdown(n int) string {
    str := ""
    for i:=n; i>0; i-=2 {
        str += fmt.Sprintf("%d", i) + ", "
    }
    str += "0!"
    return str
}
