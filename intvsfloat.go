package sprint

func IntVsFloat(i int, f float32) string {
   i2f := float32(i)
   str := "initialize variable"
   if ( i2f > f ) {
       str = "Integer"
   }
   if ( i2f < f ) {
       str = "Float"
   }
   if ( i2f == f ) {
       str = "Same"
   }
   
   return str
}

/*
package main

import "fmt"

func main () {
   i := 4
   f := float32(4.0)
   i2f := float32(i)

   fmt.Printf("i type %T", i)
   fmt.Printf("i2f type %T", i2f)
   fmt.Printf("f type %T", f)

   str := "Before if"
   if ( i2f > f ) {
       str = "Integer"
   }
   if ( i2f < f ) {
       str = "Float"
   }
   if ( i2f == f ) {
       str = "Same"
   }

   fmt.Printf(str)

}*/
