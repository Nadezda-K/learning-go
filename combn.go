//package sprint

package main
import "fmt"


func CombN(n int) []string {
	var result []string

	var Recursive func(int, string)
	Recursive = func(start int, str string) {
		if len(str) == n {
			result = append(result, string(str))
			//str = ""
			return
		}

		for i:=start; i<=9; i++ {
			fmt.Println(str)
			Recursive( i+1, str + string( rune('0' + i) ) )
		}

	}

	Recursive(0, "")

	return result
}


func main(){
    s := CombN(5)
    fmt.Println(s)

}

