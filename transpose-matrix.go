package sprint

//package main

//import "fmt"


func TransposeMatrix(matrix [][]int) [][]int {
	var transposed [][]int
	
	if len(matrix) == 0 {
		return transposed
	}
	
	rows := len(matrix)
	columns := len(matrix[0])


	var new_row []int
	for j := 0; j < columns; j++ {
		new_row = nil
		for i := 0; i < rows; i++ {
			new_row = append(new_row, matrix[i][j])
		}
		transposed = append(transposed, [][]int{new_row}...)
	}

	return transposed
}



/*
func main() {
	s := TransposeMatrix([][]int{{1, 2, 3}, {4, 5, 6}})
	fmt.Println(s)
}
*/