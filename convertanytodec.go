package sprint

//package main
import (
//		"fmt"
		"math"
		"strings"
)


func ConvertAnyToDec(s string, base string) int {	
	str := []rune(base)
	number := []rune(s)

	//check base is valid
	// if ValidBase return false, then return 0
	is_valid := ValidBase(base)
	if !is_valid {
		return 0
	}

	//check number is valid
	// if ValidNumber return false, then return 0
	is_number := ValidNumber(s, base)
	if !is_number {
		return 0
	}

	// if number is negative, convert it positive
	// keep the sign
	negative := false
	if number[0] == '-' {
		number = number[1:]
		negative = true
	}

	nbase := len(str)
	result := int(0)
	coef := int(0)

	for i := 0; i < len(number); i++ {
		coef = strings.Index( base, string(number[len(number) - 1 - i]) )
		result += int( math.Pow( float64(nbase) , float64(i) ) ) * coef
		// x = x + coef * n_base^i
	}
	
	if negative {
		result = -result
	}
	return result
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
//   Function to check if number is valid
//--------------------------------------------

func ValidNumber(nbr string, base string) bool {
	rnbr := []rune(nbr)
	var in_base bool
	
	for i, v := range rnbr {
		if i == 0 && (v == '+' || v == '-') {
			continue
		}
		in_base = strings.ContainsRune(base, v)
		if !in_base {
			return in_base
		}
	}
	return in_base
}














//--------------------------------------------
//   MAIN
//--------------------------------------------
/*func main() {
	var s int
	s = ConvertAnyToDec("92", "0123456789")
	fmt.Println(s)
	s = ConvertAnyToDec("1011100", "01")
	fmt.Println(s)
	s = ConvertAnyToDec("5C", "0123456789ABCDEF")
	fmt.Println(s)
	s = ConvertAnyToDec("did", "coding")
	fmt.Println(s)
	s = ConvertAnyToDec("-bbbac", "abc")
	fmt.Println(s)
	s = ConvertAnyToDec("bbbac", "-abc")
	fmt.Println(s)
	s = ConvertAnyToDec("-ddbac", "abc")
	fmt.Println(s)
}*/