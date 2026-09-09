package sprint

func Season(month string) string {
    str := "initializing"
    switch {
          case "jan", "feb", "dec":
                str = "winter"
          case "mar", "apr", "may":
                str = "spring"
          case "jun", "jul", "aug":
                str = "summer"
          case "sep", "oct", "nov":
                str = "autumn"
          default:
                str = "invalid input: " + month
     }
    return str
}
