package sprint

func AlphabetMastery(n int) string {
    str := ""
    for i:=1; i<=n; i++ {
        str = str + string(rune(i + 96))
    }
    return str
}
