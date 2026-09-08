package sprint

func TimeConverter(time int) {
   time := 7384
   sec := time%60
   hour := int(time/3600)
   time = time - (3600 * hour)
   min := int(time/60)
   return int(hour), int(min, int(sec)

}


//package main
//
//import "fmt"
//
//func main() {
//   time := 7384
//   sec := time%60
//   hour := int(time/3600)
//   time = time - (3600 * hour)
//   min := int(time/60)
//   fmt.Println(hour, min, sec)
//}
