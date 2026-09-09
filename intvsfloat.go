package sprint

func IntVsFloat(i int, f float32) string {
   i2f := float32(i)
   if ( i2f > f ) {
       str := "Integer"
   }
   if ( i2f < f ) {
       str := "Float"
   }
   if ( i2f == f ) {
       str := "Same"
   }
   
   return str
}
