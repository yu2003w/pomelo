package example

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

type Niler interface {
    Ni() 
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f) 
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    if v == nil {
        fmt.Println("nil")
        return float64(-1)
    }
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Ginter() {
    var a Abser
    f := Myfloat(-math.Sqrt2)
    v := Vertex{3,4}
    a = f
    fmt.Println("Interface testing")
    fmt.Printf("(%v, %T), real value = %f\n", a, a, a.Abs())

    a = &v
    fmt.Printf("(%v, %T), real value = %f\n", a, a, a.Abs())

    var n *Vertex
    a = n
    fmt.Printf("(%v, %T), real value = %f\n", a, a, a.Abs())
    
    var ni Niler
    fmt.Printf("(%v, %T)\n", ni, ni)
    
    // runtime error for below instruction
    // ni.Ni()
    
}

func desc(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

func (v Vertex) String() string {
    return fmt.Sprintf("Vertex{%v, %v}", v.X, v.Y)
}

func Einter() {
    var ei interface{}
    ei = 42
    desc(ei)
    ei = "empty interface"
    desc(ei)
    ei = Vertex{3,4}
    desc(ei)
    t, ret := ei.(int)
    if ret {
        fmt.Println(t)
    }else{
        fmt.Println("Value is not type int")
    }

    switch v := ei.(type) {
    case int:
        fmt.Printf("int value is %d\n", v)
    case string:
        fmt.Printf("string value is %q\n", v)
    case Vertex:
        fmt.Println("Vertex value is", v)
    default:
        fmt.Println("Unknown type")
    }
    
}

