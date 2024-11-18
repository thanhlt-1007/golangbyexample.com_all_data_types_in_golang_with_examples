package main

import (
    "fmt"
    // "math/bits"
    "reflect"
    "unsafe"
)

func main() {
    // int
    // sizeOfIntInbits := bits.UintSize
    // fmt.Printf("%d bits\n", sizeOfIntInbits)

    // var a int
    // fmt.Printf("%d bytes \n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // b := 2
    // fmt.Printf("b's type is %s\n", reflect.TypeOf((b)))

    // int8
    // var a int8 = 2
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // int16
    // var a int16 = 2
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // int32
    // var a int32 = 2
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // int64
    var a int64 = 2
    fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
