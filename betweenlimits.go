package sprint

func BetweenLimits (from, to rune) string {
    var ext_rune rune
    var acc_str string

    if ( from > to ) {
      ext_rune = from
      from = to
      to = ext_rune
    }

    for i:=from+1; i<to; i++ {
        acc_str = acc_str + string(i)
    }
    
    return acc_str
}

/*
package main

import "fmt"

func main() {
    from := 'j'
    to := 'f'
    
    var ext_rune rune
    var acc_str string

    if ( from > to ) {
      ext_rune = from
      from = to
      to = ext_rune
    }

    for i:=from+1; i<to; i++ {
        acc_str = acc_str + string(i)
    }
    fmt.Printf(acc_str)
}
*/
