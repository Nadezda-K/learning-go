package sprint

import "math"

func PointDistance(x1, y1, x2, y2 float64) float64 {
    d := math.Sqrt( math.Pow(x2-x1,2) + math.Pow(y2-y1,2) )
    return d
}

