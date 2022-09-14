package main

import "fmt"

func main() {
    ids := []int{55,44,33,6,23}

    // Loop through ids
    for i, id := range ids {
        fmt.Printf("%d -- ID: %d\n", i, id)
    }

    // Without index
    for _, id := range ids {
        fmt.Printf("ID: %d\n", id)
    }

    // Add ids together
    sum := 0
    for _, id := range ids {
        sum += id
    }
    fmt.Println("Sum", sum)

    // Range with map
    emails := map[string]string{"Alex":"alex@gmail.com", "Ken":"ken@gmail.com"}

    for k, v := range emails {
        fmt.Printf("%s: %s\n", k, v)
    }

    for k := range emails {
        fmt.Println("Name: " + k)
    }
}
