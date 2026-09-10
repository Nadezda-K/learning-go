package sprint

func ReverseAlphabet(step int) string {
    if ( step <= 0 ) {
       step = 1
    }
    
    str = "="
    for i:=26; i>0; i-=step {
        str = str + string( rune(i + 96) )
    }
    return str
}
