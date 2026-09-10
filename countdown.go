package sprint

func Countdown(n int) {
    str := ""
    for i:=n; i>0; i-=2 {
        str += i + ", "
    }
    str += "0!"
}
