package main

import "fmt"

func main() {
    //Arrays
    // var fruitArr [2]string

    // Assign values
    // fruitArr[0] = "Apple"
    // fruitArr[1] = "Banana"

    // Declare and Assign simultaneously
    // fruitArr := [2]string{"Apple","Banana"}

    // fmt.Println(fruitArr)
    // fmt.Println(fruitArr[1])

    fruitSlice := []string{"Apple","Banana","Orange"}

    fmt.Println(fruitSlice)
    fmt.Println(len(fruitSlice))
    fmt.Println(fruitSlice[1:2]) // --> [Banana] Inclusive:Exclusive
}
