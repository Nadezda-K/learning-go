package sprint

//package main
//import "fmt"


func NbrBase(n int, base string) string {	
	str := []rune(base)	

	//check base is valid
	// if ValidBase return false, then return "NV"
	is_valid := ValidBase(base)
	if !is_valid {
		return "NV"
	}

	// if number is negative, convert it positive
	// keep the sign
	sign := ""
	if n < 0 {
		n = -n
		sign = "-"
	}

	var result []rune
	nbase := len(str)
	var digit int

	for n > 0 {
		digit = n % nbase
		result = append( result, rune(digit+'0') )
		n = n / nbase
		//fmt.Println(digit, string(digit+'0'))
	}

	number := ""
	for _,v := range result {
		//fmt.Println(v, string(v), string(base[v-'0']))
		number = string(base[v-'0']) + number
	}

	return sign + string(number)
}

//--------------------------------------------
//   Function to check if base is valid
//--------------------------------------------
func ValidBase(base string) bool {	
	str := []rune(base)	
	//check if less then  2 characters in base
	if len(str) < 2 {
		return false
	}

	//check if there are duplicates in base
	for i,ch := range str {
		if ch == '+' || ch == '-' {
			return false
		}
		for j,v := range str {
			if ch == v && i != j {
				return false
			}
		}
	}
	return true
}


//--------------------------------------------
//   MAIN
//--------------------------------------------
/*func main() {
	var s string
	s = NbrBase(92, "0123456789")
	fmt.Println(s)
	s = NbrBase(-92, "01")
	fmt.Println(s)
	s = NbrBase(92, "0123456789ABCDEF")
	fmt.Println(s)
	s = NbrBase(-91, "coding")
	fmt.Println(s)
	s = NbrBase(92, "abb")
	fmt.Println(s)
	s = NbrBase(92, "ab+")
	fmt.Println(s)
	s = NbrBase(92, "-bcde")
	fmt.Println(s)
	s = NbrBase(92, "e")
	fmt.Println(s)
}*/