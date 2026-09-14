package sprint

//package main
//import "fmt"

func GenerateRange(min, max int) []int {
    var slice []int
    if ( min < max ) {
        length := max - min
        slice = make([]int, 0, length)
        for i:=min; i<max; i++ {
            slice = append(slice, i)
        }
    }
    return slice
}

/*
func main(){
    s := GenerateRange(3,-9) 
    fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
*/