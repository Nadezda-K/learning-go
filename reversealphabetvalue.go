package sprint

func ReverseAlphabetValue(ch rune) rune {
    return rune( int('z')-(int(ch)-int('a')) )
}

