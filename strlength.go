package sprint

//package main
//
//import "fmt"

func StrLength(s string) []int  {
	return []int{ len([]rune{s}), len(s)}
}

/*
func main() {
	s := StrLength("kood")
	fmt.Println(s)
}
*/