//package sprint

package main
import "fmt"


func CombN(n int) []string {
	var result []string

    var MakeNumbers func(string, int)
	MakeNumbers = func(number string, start int) {
		if len(number) == n {
			result = append(result, number)
			return
		}
		for i:=start; i<=9; i++ {
			MakeNumbers(number + string('0'+i), i+1)
		}
	}

	MakeNumbers("", 0)
	return result
}


func main(){
    s := CombN(4)
    fmt.Println(s)

}
