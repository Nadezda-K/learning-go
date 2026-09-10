package sprint

func Countdown(n int) string {
    str := ""
    for i:=n; i>0; i-=2 {
        str += string(rune(i+48)) + ", "
    }
    str += "0!"
    return str
}
