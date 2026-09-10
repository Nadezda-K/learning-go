package sprint

func AlphabetMastery(n int) string {
    str := ""
    for i:=0; i<=n; i++ {
        str = str + string(rune(i + 97))
    }
    return str
}
