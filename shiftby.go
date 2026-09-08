package sprint

func ShiftBy(r rune, step int) rune {
    return (int(r) - int('a') + step) % 26 + int('a')
}

