package sprint

//package main
//
//import "fmt"

func ToUpperCase(s string) string {
	new_str := ""
	for _, v := range s {
		if v >= rune('a') && v <= rune('z') {
			v = v - 'a' + 'A' 
		}
		new_str += string(v)
	}
	return new_str
}

/*
func main() {
	s := ToUpperCase("Hello! How's your day going?")
	fmt.Println(s)
}
*/