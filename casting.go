package sprint

import "math"

func Casting(n float64) int {
    return n - math.Mod(n,1)
}
