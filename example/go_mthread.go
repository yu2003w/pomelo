package example

import (
    "fmt"
    "time"
)

func run(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100*time.Millisecond)
        fmt.Println(s)
    }
}

func Scon() {
    go run("Sub thread")
    run ("main thread")
}

