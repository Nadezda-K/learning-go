package sprint

//package main
//import "fmt"

func make_divisors_list(x int) []int {
    // function decompose number into its prime factors
    
    var list []int
    
    for x > 1 {
        for i:=2; i<=x; i++ {
            if ( x%i == 0 ) {
                list = append( list, i )
                x = x/i
                
                break // Only breaks the inner loop
            }
        }
    }
    return list
}



func is_in_slice(val int, list []int) (bool, int) {
    // function check if val is in list
    // return true if in the list

    for i, v := range list {
        if ( val == v ) {
            return true, i
        }       
    }
    return false, -1
}



func LCM(a, b int) int {
    if ( a == 0 ) || ( b == 0 ) { 
        return 0
    }
    
    list_a := make_divisors_list(a)
    list_b := make_divisors_list(b)     
    
    if ( len(list_a) > len(list_b) ) {
        //swap lists values
        list_a, list_b = list_b, list_a
    }

    // Compare elements of the smaller list with those of the larger list
    // and append only the elements that do not exist in the larger list. 
    
    for _, v := range list_a {
        check, j := is_in_slice(v, list_b)
        if check == true {
           list_b = append( list_b[:j], list_b[j+1:]...)
           // ... is important for appending a list not a single value
        }        
    }
    
    result_list := append(list_a, list_b...)
    
    lcm := int(1)
    for _, v := range result_list {
        lcm *= v
    }
    
    return lcm
}

/*
func main() {
    fmt.Println("Least Common Multiple = ", LCM(30, 18))
}*/
