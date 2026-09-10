package sprint

//imprt "fmt"

func Pairs() string {
    str := ""
    var str_i, str_k string
    for i:=0; i<100; i++ {
        if ( i < 10 ) {
            str = str + "0"
        }
        str_i = fmt.Sprintf("%d", i)
        str = str + str_i + " "
        
        for k:=i+1; k<100; k++ {
            if ( k < 10 ) {
                str = str + "0"
            }
            str_k = fmt.Sprintf("%d", k)
            str = str + str_k + ", "
        }
    }
    
    return str
}
