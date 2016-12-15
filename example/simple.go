package example

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

var pkgname string

var auth string = "Jared"

func Simple() {
    pkgname = "example"
    fmt.Println("Hello Go!")
    fmt.Println("Time is ", time.Now())
    fmt.Println("My favourate number is ", rand.Intn(10))
    fmt.Println("Pi value is ", math.Pi)
    fmt.Println(Add(4,5))
    fmt.Println(swapStr(" hello ", " Go! "))
    x,y := split(17)
    fmt.Println("x = ", x, ", y = ", y)
    fmt.Println("author is ", auth)
    
    v := 42
    fmt.Printf("Type of v is %T\n", v)

    const hello = "hello go!"
    fmt.Println(hello)

}

func Add(x int, y int) int {
    fmt.Println("function add is called in package ", pkgname)
    return x + y
}

func swapStr(x, y string) (string, string) {
    fmt.Println("function swapStr is called in package ", pkgname)
    return y, x
}

func split(a int) (x, y int) {
    fmt.Println("function split is called in package ", pkgname)
    x = a/14*9
    y = a - x
    return
}

