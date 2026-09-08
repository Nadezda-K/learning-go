package sprint

import "math"

func Casting(n float64) int {
    return n - math.Remainder(n,1)
}
