package example

import (
    "fmt"
)

type Pos struct {
    X int
    Y int
}

func Gpointer() {
    i, j := 20, 42
    p := &i

    fmt.Printf("\n*p is %d\n", *p)
    *p = 21
    fmt.Printf("*p changed to %d\n", i)
   
    p = &j
    fmt.Printf("*p is %d\n", *p)
    *p = *p/7
    fmt.Printf("*p is %d\n", j)
}

func Garr() {
    var data [5] string
    data[0] = "Arrary"
    data[1] = "Strings"
    data[2] = "go"
    data[3] = "languages"
    data[4] = "type"

    fmt.Println(data)
    fmt.Println(data[0], data[2])
   
    a := data[2:4]
    a[0] = "type"
    a[1] = "go"
    fmt.Println(data)

    b := data[4:]
    b[0] = "language"
    fmt.Println(data)

    fmt.Printf("slice length %d, array length %d\n", len(a), cap(a))
}

func Gslice() {
    a := make([]int, 5)
    fmt.Println(a)

    b := make([]int, 3, 5)
    fmt.Println(b)
    fmt.Printf("slice length %d, arrary length %d\n", len(b), cap(b))

    b = append(b, 2, 3, 4)
    fmt.Println(b[3:6])
}

func Gmap() {
    data := make(map[string]Pos)
    data["bell"] = Pos{2,4}
    data["labs"] = Pos{3,6}
    fmt.Println(data)

    e, ok := data["lu"]
    if ok == true {
        fmt.Println(e)
    } else {
        fmt.Println("key not found in map")
    }

    e, ok = data["bell"]
    if ok == true {
        fmt.Println(e)
    } else {
        fmt.Println("key not found in map")
    }
}

func adder() func(int) int {
    sum := 0;
    return func(x int) int {
        sum += x
        return sum
    }
}

func Gclos() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}

