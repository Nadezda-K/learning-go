package sprint

//package main

//import "fmt"

func StrSplitBy(s, sep string) []string {
	var arr []string
	pos := int(1) 

	if s == "" || sep == "" {
		return arr
	}



	for i,j :=0, len(sep); j <= len(s); i,j = i+1, j+1 {// i=0, j=size (=3 for "YOU");  
														// i= i+1, j=j+1
		window := s[i:j]
		fmt.Println(window)
		if window == sep {
			arr = append(arr, s[pos-1:i])
			pos = j+1
		}		
	}
	arr = append(arr, s[pos-1:])
	return arr
}

/*
func main() {
	s := StrSplitBy("YOUHowYOUhaveYOUyouYOUbeen?", "YOU")

	fmt.Println(s)
}
*/