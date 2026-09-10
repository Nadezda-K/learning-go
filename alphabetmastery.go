package sprint

func AlphabetMastery(n int) string {
    var str string
    for i:=0; i<=n; i++ {
        str = str + string(rune(i))
    }
    return str
}
