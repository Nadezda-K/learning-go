//package sprint

package main
import (
//		"fmt"
		"math"
		"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	//base_from := []rune(base_from)
	//base_to := []rune(base_to)
	number := []rune(nbr)

	//check base is valid
	// if ValidBase return false, then return 0
	is_valid := ValidBase(baseTo)
	if !is_valid {
		return "NV"
	}
	is_valid = ValidBase(baseFrom)
	if !is_valid {
		return"NV"
	}

	//check number is valid
	// if ValidNumber return false, then return 0
	is_number := ValidNumber(nbr, baseFrom)
	if !is_number {
		return "NV"
	}

	// if number is negative, convert it positive
	// keep the sign
	negative := false
	if number[0] == '-' {
		number = number[1:]
		negative = true
	}
	
	n := ConvertAnyToDec(nbr, baseFrom)
	result := NbrBase(n, baseTo) 


	if negative {
		result = "-" + string(result)
	}
	return result

}


//--------------------------------------------
//   Function to covert from any to decimal
//--------------------------------------------

func ConvertAnyToDec(s string, base string) int {	
	str := []rune(base)
	number := []rune(s)

	nbase := len(str)
	result := int(0)
	coef := int(0)

	for i := 0; i < len(number); i++ {
		coef = strings.Index( base, string(number[len(number) - 1 - i]) )
		result += int( math.Pow( float64(nbase) , float64(i) ) ) * coef
		// x = x + coef * n_base^i
	}
	
	return result
}


//--------------------------------------------
//   Function to covert from decimal to any
//--------------------------------------------

func NbrBase(n int, base string) string {	
	str := []rune(base)	

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

	return string(number)
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
	var s string
	s = ConvertAnyToAny("100001", "01", "0123456789")
	fmt.Println(s)
}*/