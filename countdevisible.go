package sprint

func CountDivisible(from, to, step, divisor int) int {
    if ( step <= 0 ) || ( devisor == 0 ) {
        return 0
    }

    count := 0
    for i:=from; i<to; i++ {
        if (i % step == 0 ) {
            if (i % devisor == 0) {
                count = count + 1
            }
        }
    }
}
