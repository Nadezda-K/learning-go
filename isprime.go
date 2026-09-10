package sprint

func IsPrime(n int) bool {
    if ( n <= 2) {
        return false
    }
    for i:=3; i<n; i++ {
       if ( n%i == 0) {
           return false
       }
    }
    return true
}

