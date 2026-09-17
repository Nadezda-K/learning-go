package sprint

//package main
//import "fmt"

type Point struct {
	X float32
	Y float32
	Text string 
}

func MakePoint(x, y float32, text string) Point {
	var point Point

	point.X = x
	point.Y = y
	point.Text = text

	return point
}

/*
func main() {
    var s int

	fmt.Println(s)
}
*/
