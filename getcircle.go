package sprint

//package main
//import "fmt"

type Circle struct {
	Radius float32
	Diameter float32
	Area float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	var c Circle
	pi := float32(3.14)
	c.Radius = r 
	c.Diameter = 2 * r
	c.Area = pi * r *r
	c.Perimeter = 2 *pi * r
}

/*
//func main() {
    var s Circle
	s = GetCircle(5)

	fmt.Println(s.Radius, s.Diameter, s.Area, s.Perimeter)
}
*/
