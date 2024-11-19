package main

import (
    "fmt"
    // "math/bits"
    "reflect"
    "unsafe"
)

type sample struct {
    a int
    b string
}

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
    // var a int64 = 2
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // uint
    // var a uint
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // uintptr
    // s := &sample{a: 1, b: "test"}
    // p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))
    // fmt.Println((*(*string)(p)))

    // uint8
    // var a uint8
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // uint16
    // var a uint16
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // uint32
    // var a uint32
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // uint64
    // var a uint64
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // float32
    // var a float32
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // float64
    // var a float64
    // fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    // fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    // byte (unit8)
    var r byte = 'a'
    fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
    fmt.Printf("Type: %s\n", reflect.TypeOf(r))
}
