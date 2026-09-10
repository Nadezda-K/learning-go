package sprint

//imprt "fmt"

func Pairs() string {
    str := ""
    for i:=0; i<100: i++ {
        if ( i < 10 ) {
            str = str + "0"
        }
        str = str + string(i) + " "
        for k:=i+1; k<100; k++ {
            if ( k < 10 ) {
                str = str + "0"
            }
            str = str + string(k) + ", "
        }
    }
    return str
}
