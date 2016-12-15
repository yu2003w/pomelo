package main

import (
    "fmt"
    "runtime"
    "github.com/yu2003w/pomelo/example"
)

func main() {
    fmt.Println("Hello pomelo")

    example.Simple()
    example.Constnum()

    fmt.Println("Go runs on ", runtime.GOOS)
    example.Gettoday()
    example.Defergame()

    example.Gpointer()
 
    // struct type
    v := example.Pos{1,2}
    fmt.Printf("Pos(%d, %d)\n", v.X, v.Y)
    v.X = 4
    v.Y = 8
    fmt.Printf("Pos(%d, %d)\n", v.X, v.Y)
    p := &v
    p.X = 5
    fmt.Println("Updated pos -> ", v)
   
    example.Garr()
    example.Gslice()
    example.Gmap()
    example.Gclos()
    example.Ginter()
    example.Einter()
    example.Scon()
    example.Gchan()
    example.Gsel()
}

