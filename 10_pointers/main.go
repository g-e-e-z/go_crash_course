package main

import "fmt"

func main() {
    a := 5
    b := &a

    fmt.Println(a, b)
    fmt.Println(a, &a)

    fmt.Printf("%T\n", a)
    fmt.Printf("%T\n", b)

    // Use * to read val from address
    fmt.Println(*b)
    // fmt.Println(*&a)

    // Change val with pointer
    *b = 10
    fmt.Println(a)

}
