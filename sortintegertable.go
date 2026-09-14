package sprint

//package main
//import "fmt"

func SortIntegerTable(table []int) []int {
    for i:=0; i<len(table); i++ {
        for j:=0; j<len(table)-i-1; j++ {
            if table[j] > table[j+1] {
                table[j], table[j+1] = table[j+1], table[j]
            }
        }        
    }    
    return table 
}

/*
func main(){
    s := SortIntegerTable([]int{2, 0, 5, 4, 1, 3})
    fmt.Println(s)

}
*/