package sprint

func Accumulate(n int) int {
    if ( n >= 0 ) {
       acc := int(0)
       for i:=0; i<=n; i++ {
           acc = acc + i
       }
    }
    if ( n < 0 ) {
        acc = 0
    }

    return acc
}
