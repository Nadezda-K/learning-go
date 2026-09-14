package sprint

//package main
//import "fmt"

func BalanceOut(arr []bool) []bool {
    //var slice []bool
    count_t := 0
    count_f := 0
    for i,v := range arr {
        fmt.Println(i, v)
        if v {
            count_t++
        } else {
            count_f++
        }        
    }
    fmt.Println(count_t, count_f)
    if count_f > count_t {
        for i:=count_t; i<count_f; i++ {
            arr = append(arr, true )
        }
    } else {
        for i:=count_f; i<count_t; i++ {
            arr = append(arr, false )
        }
    }
    return arr
}

/*
func main(){
    s := BalanceOut([]bool{true, true, true, false, false, false})
    fmt.Println(s)

}
*/