package sprint

//package main
//import "fmt"

func FilterBySum(arr [][]int, limit int) [][]int {
    var new_arr [][]int
    var sum int
    for i:=0; i<len(arr); i++ {
        sum = 0
        for j:=0; j<len(arr[i]); j++ {
            sum += arr[i][j]
        }
        if sum >= limit {
            new_arr = append(new_arr, [][]int{arr[i]}...)
        }
    }
    return new_arr
}

/*
func main(){
    s := FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9)
    fmt.Println(s)

}
*/