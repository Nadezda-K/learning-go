package sprint

func Fibonacci(n int) int {
    if ( n < 0 ) {
        return 0
    }
    if  ( n == 1 ) {
        return 1
    }
    
    ni2 := int(0)
    ni1 : int(1)
    var ni int

    for i:=2; i<=n; i++ {
        ni = ni2 + ni1
        ni2 = ni1
        ni1 = ni
    }
    return ni
}
