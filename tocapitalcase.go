package sprint

//package main
//
//import "fmt"

func ToCapitalCase(s string) string {
	new_str := ""
	previous_is_letter := bool(false)
	for _, v := range s {

		//fmt.Println(previous_is_letter, string(v))

		if (v < rune('a') || v > rune('z')) && (v < rune('A') || v > rune('Z')) {
			previous_is_letter = false
		} else {
			if v >= rune('a') && v <= rune('z') {
				if previous_is_letter == false {
					v = v - 'a' + 'A'
				}
			}
			if v >= rune('A') && v <= rune('Z') {
				if previous_is_letter == true {
					v = v - 'A' + 'a'
				}
			}
			previous_is_letter = true
		}
		
		new_str += string(v)
		//fmt.Println(string(v), previous_is_letter)
	}
	return new_str
}

/*
func main() {
	s := ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?")
	fmt.Println(s)
}
*/