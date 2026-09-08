package sprint

func TimeConverter(time int) (int, int, int) {
   sec := time%60
   hour := int(time/3600)
   time = time - (3600 * hour)
   min := int(time/60)
   
   return hour, min, sec
}

