package sprint
//package main

import "fmt"

func Pairs() string {
// func main()
    str := ""
    var str_i string
    var str_k string
    for i:=0; i<100; i++ {
        str_i = ""
        if ( i < 10 ) {
            str_i = "0" 
        }
        str_i = str_i + fmt.Sprintf("%d", i)
        
        for k:=i+1; k<100; k++ {
            str_k = ""
            if ( k < 10 ) {
                str_k = "0"
            }
            str_k = str_k + fmt.Sprintf("%d", k)
            str = str + str_i + " " + str_k
            if ( i != 98 ) || ( k != 99 ) {
                str = str + ", "
            }
        }
    }
    
    return str
//    fmt.Println(str)
}
