package sprint

//package main
//
//import "fmt"

func StrReverse(s string) string  {
	var reverse string 
	r_str := []rune(s)
	for i:=len(r_str)-1; i>=0; i-- {
		reverse += string(r_str[i]) 	
	}
	return reverse
}

/*
func main() {
	s := StrReverse("kood")
	fmt.Println(s)
}
*/