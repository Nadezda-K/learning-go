package sprint

//package main
//
//import "fmt"

func SubstrIndex(s string, toFind string) int {
	if s == "" ||  toFind== "" {
		return -1
	}

	for i,j := 0,len(toFind); j<=len(s); i,j = i+1, j+1 {
		if s[i:j] == toFind {
			return i
		} 
	}
	return -1	
}

/*
func main() {
	//s := SubstrIndex("How are you?", "o")
	//s := SubstrIndex("How are you doing?", "ou")
	s := SubstrIndex("You can do it!", " od")
	fmt.Println(s)
}
*/