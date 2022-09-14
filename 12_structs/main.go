package main

import (
    "fmt"
    "strconv"
)

// Define `Person` struct
// type Person struct {
//     firstName string
//     lastName string
//     city string
//     gender string
//     age int
// }

type Person struct {
    firstName, lastName, city, gender string
    age int
}

// Two types of methods on structs
// 1. Value receivers -> Just do calculations, dont change anything
// 2. Pointer receivers -> Actually changing something
// Do not go inside the struct

// Greeting Method
// func (identifier, nameOfStruct)
func (p Person) greet() string {
    return "Hello " + p.firstName + " " + p.lastName + ". Age: " + strconv.Itoa(p.age)
}

// Birthday Method, does not return anything
func(p *Person) hasBirthday() {
    p.age ++
}

// Married
func(p *Person) getMarried(name string) {
    if p.gender == "f" {
        p.lastName = name
    }
}

func main() {
    // Init person using struct
    person1 := Person{
        firstName: "Alex",
        lastName: "Gordon",
        city: "Denver",
        gender: "f",
        age: 27,
    }
    // person1 := Person{"Alex","Gordon","Denver","m",27}

    // fmt.Println(person1)
    // fmt.Println(person1.age)
    // fmt.Println(person1.age)

    fmt.Println(person1.greet())
    person1.hasBirthday()
    fmt.Println(person1.greet())

    person1.getMarried("Gugenheim")
    fmt.Println(person1.greet())

}
