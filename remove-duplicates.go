//package sprint

package main

import "fmt"


func RemoveDuplicates(arr []int) []int {
	var new []int
	var check bool
	
	if arr == nil {
		return new
	}
	
	new = append(new, arr[0])
	for _, b := range arr {
		check = false
		for _, a := range new {
			if a == b {
				check = true
				break
			}
		}
		if check == false {
			new = append(new, b)
		}
	} 

	return new
}




func main() {
	s := RemoveDuplicates([]int{1, 2, 3, 2, 4, 8, 8, 1, 2, 0, 8})
	fmt.Println(s)
}
