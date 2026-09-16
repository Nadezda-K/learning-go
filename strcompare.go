package sprint

//package main
//
//import "fmt"

func StrCompare(a, b string) int {
	var str []rune
	a_rune := []rune(a)
	b_rune := []rune(b)
	if len(a_rune) <= len(b_rune) {
		str = a_rune
	} else {
		str = b_rune
	}

	for i, _ := range str {
		if a_rune[i] > b_rune[i] {
			return 1
		}
		if a_rune[i] < b_rune[i] {
			return -1
		}	
	}

	if len(a_rune) > len(b_rune) {
		return 1
	} else {
		return 0
	}
}

/*
func main() {
	//s := StrCompare("Hi!", "Hi!")
	//s := StrCompare("Day", "ay")
	s := StrCompare("weekday", "week")
	fmt.Println(s)
}
	*/
