package sprint

//package main
import "fmt"

type Point struct {
	X float32
	Y float32
	Text string 
}

func PointText(p Point) Point {
	p.Text = fmt.Sprintf( "Text at (%f, %f)", p.X, p.Y )
	return p
}

/*
func main() {
    var s Point
	s.X = 1.5
	s.Y = 0.93784
	s.Text = "some text"
	fmt.Println(s.Text)

	s = PointText(s)
	fmt.Println(s.Text)
}
*/
