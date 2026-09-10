package sprint

func IsPrime(n int) bool {
    for i:=1; i<n; i++ {
       if ( n%i == 0) {
           return false
       }
    }
    return true
}

