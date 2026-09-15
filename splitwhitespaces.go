package sprint

//package main
//
//import "fmt"

func SplitWhitespaces(s string) []string {
	var arr []string
	word := ""
	for _, v := range s {
		if v == ' '|| v == '\n' || v == '\t' {
			arr = append(arr, word)
			word = ""
		} else {
			word += string(v)
		}
	}

	arr = append(arr, word)
	return arr
}

/*
func main() {
	s := SplitWhitespaces("Hello!	How have you	been?")
	fmt.Println(s)
}
*/