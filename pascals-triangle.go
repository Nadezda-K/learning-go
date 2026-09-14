package sprint

//package main
//import "fmt"


func PascalsTriangle(n int) [][]int {
	var triangle [][]int

	row := []int{1}
	triangle = append(triangle, [][]int{row}...)

	var nj int
	for i:=1; i<n; i++ {
		nj = 0
		row = []int{}
		for j:=0; j<=i; j++ {
			if len(triangle[i-1]) == j{
				row = append(row, 1)
			} else {
				row = append(row, nj+triangle[i-1][j])
				nj = triangle[i-1][j]
				}
		}
		triangle = append(triangle, [][]int{row}...)
	}
	return triangle
}

/*
func main(){
    s := PascalsTriangle(5)
    fmt.Println(s)

}
*/