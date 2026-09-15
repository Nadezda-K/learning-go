package sprint

//package main
//
//import "fmt"

func IsLower(s string) bool {
	check := bool(true)
	for _, v := range s {
		if v < 97 || v> 122 {
			check = false
		}
	}
	return check
}

/*
func main() {
	s := IsLower("kood!")
	fmt.Println(s)
}
*/