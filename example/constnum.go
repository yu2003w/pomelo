package example

import "fmt"

const (
    BIG = 1 << 100
    SMALL = 1 >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x*0.1
}

func Constnum() {
    fmt.Println(needInt(SMALL))
    fmt.Println(needFloat(BIG))
    fmt.Println(needFloat(SMALL))
}

