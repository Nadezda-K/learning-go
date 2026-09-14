//package sprint

package main
import "fmt"

func StrToInt(s string) int {
    result := int(0)
    sign := int(1)
    beginning_is_zero := true

    for i:=0; i<len(s); i++ {
        if ( rune(s[i]) < rune('0') ) || ( rune(s[i]) > rune('9') ) {
            if ( i == 0 ) && ( rune(s[i]) == 43 ){
                sign = 1
            } else if ( i == 0) && ( rune(s[i]) == 45 ) {
                sign = -1
            } else {
                return 0
            }
        }
        if ( rune(s[i]) > rune('0') ) && ( beginning_is_zero == true ) {
            beginning_is_zero = false
        }
        if ( beginning_is_zero == false ) {
            result = result*10 + int(s[i]-'0')
        }
    }
    return result*sign
}



func BulkAtoi(arr []string) []int {
    var bulk []int
	
	for i:=0; i<len(arr); i++ {
		bulk = append(bulk, StrToInt(arr[i]) )
	}

	return bulk
}


func main(){
    s := BulkAtoi([]string{"8", "kood", "-13"})
    fmt.Println(s)

}
