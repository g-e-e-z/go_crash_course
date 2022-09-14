package main

import "fmt"

func main() {
    x := 10
    y := 10

    // if else
    if x <= y {
        fmt.Printf("%d is less than or equal to %d\n", x, y)
    } else {
        fmt.Printf("%d is less than %d\n", y, x)
    }

    // else if
    color := "green"

    if color == "red" {
        fmt.Println("Color is red")
    } else if color == "blue" {
        fmt.Println("Color is blue")
    } else {
        fmt.Println("Color is neither red or blue")
    }

    // switch
    switch color {
    case "red":
        fmt.Println("Color is red")
    case "blue":
        fmt.Println("Color is blue")
    default:
        fmt.Println("Color not blue or red")
    }

}
