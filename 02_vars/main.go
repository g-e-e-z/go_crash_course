package main

import "fmt"

func main() {
    // MAIN TYPES
    // sting
    // bool
    // int
    // int int8 int16 int32 int64
    // uint uint8 uint16 uint32 uint64
    // byte - alias for uint8
    // rune - alias for int32
    // float32 float64
    // complex64 complex128

    //Using var
    // var name string = "Alex"
    var age int32 = 27
    const isCool = true
    var size float32 = 196.5

    // Shorthand -> Must be within a function
    name, email := "Alex", "alex@.com"


    fmt.Println(name, age, size, email, isCool)
    fmt.Printf("%T\n", size)

}

