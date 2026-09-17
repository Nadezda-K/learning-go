package sprint

//package main
//import "fmt"

type Coords struct {
	X int
	Y int
}

type Rectangle struct {
	Width int
	Height int
	Area int
	Perimeter int
}

func GetRectangle(min, max Coords) Rectangle {
	var r Rectangle
	r.Width = max.X - min.X
	r.Height = max.Y -min.Y
	r.Area = r.Width * r.Height
	r.Perimeter = 2 * (r.Width + r.Height)

	return r
}

/*
//func main() {
    var s Circle
	s = GetCircle(5)

	fmt.Println(s.Radius, s.Diameter, s.Area, s.Perimeter)
}
*/
