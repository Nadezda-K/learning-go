package sprint

//package main
//import "fmt"

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    //var slice []float64
    if from > to {
        from, to = to, from
    }
    arr = append(arr[:from], arr[to:]...)
    return arr
}

/*
func main(){
    s := RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1) 
    fmt.Println(s, len(s), cap(s))

}
*/