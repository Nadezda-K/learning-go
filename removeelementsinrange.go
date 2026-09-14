package sprint

//package main
//import "fmt"

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
    //var slice []float64
    if from > to {
        from, to = to, from
    }
    if from <=0 && to>=len(arr) {
        return make([]float64,0,0)
    }

    arr = append(arr[:from], arr[to:]...)
    return arr
}

/*
func main(){
    s := RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 0, 7) 
    fmt.Println(s, len(s), cap(s))

}
*/