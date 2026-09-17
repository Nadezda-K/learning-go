package sprint

//package main
//import "fmt"

type Contestant struct {
	Name string
	Scores []int
}

func Total(score []int) int {
	sum := int(0)

	for _, v := range score {
		sum += v
	}

	return sum
}
func GetWinner(c1, c2 Contestant) string {
	if Total(c1.Scores) > Total(c2.Scores) {
		return c1.Name
	} else {
		return c2.Name
	}
}

/*
//func main() {
    var s Circle
	s = GetCircle(5)

	fmt.Println(s.Radius, s.Diameter, s.Area, s.Perimeter)
}
*/
