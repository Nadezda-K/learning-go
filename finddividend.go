package sprint

func FindDividend(from, to, divisor int) int {
    x := int(-1)
    for  i:=from; i<to; i++ {
       if (i%divisor == 0) {
          x = i
          return x
       }
    }
    return x
}
