package sprint

//package main
//
//import "fmt"

func StrConcatWith(strs []string, sep string) string {
	new_str := ""
	for _, v := range strs {
		if (new_str != "") {
			new_str += sep
		}
		new_str += v
	}
	return new_str
}

/*
func main() {
	toConcat := []string{"Three", " Two", " One", " Go!"}
	s := StrConcatWith(toConcat, ".")
	fmt.Println(s)
}
*/