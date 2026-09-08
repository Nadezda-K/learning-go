package sprint

func RuneChecksum(a, b rune) rune {
   my_int1 := int(my_rune1) - int('a') + 1
   my_int2 := int(my_rune2) - int('a') + 1
   new_int := (my_int1 * my_int2) % 26
   
   return rune(new_int)
}


//package main
//
//import "fmt"
//
//func main () {
//    my_rune1 := 'c'
//    my_rune2 := 'd'
//
//   my_int1 := int(my_rune1) - int('a') + 1
//   my_int2 := int(my_rune2) - int('a') + 1
//   new_int := (my_int1 * my_int2) % 26
//   new_rune := rune(new_int)
//   fmt.Println(new_rune, string(new_rune))
//
//}
