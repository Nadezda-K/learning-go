package sprint
//package main

import "fmt"

func Combinations() string {
//func main() {
   str := ""
   
   for i:=0; i<8; i++ {
       for j:=0; j<9; j++ {
           for k:=0; k<10; k++ {
               if (j > i) && ( k > j ) {
                   str += fmt.Sprintf("%d%d%d", i, j, k)
                   if (i != 7) || (j != 8) || (k != 9) {
                       str += ", "
                   }
               }
           }
       }
   }
    return str
//    fmt.Println(str)
}
