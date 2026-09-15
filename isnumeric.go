package sprint

//package main
//
//import "fmt"

func IsNumeric(s string) bool {
	check := bool(true)
	for _, v := range s {
		if v < rune('0') || v > rune('9') {
			check = false
		}
	}
	return check
}

/*
func main() {
	s := IsNumeric("1234")
	fmt.Println(s)
}
*/