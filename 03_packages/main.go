package main

import (
    "fmt"
    "math"
    "github.com/g-e-e-z/go_crash_course/03_packages/strutil"
)

func main() {
    fmt.Println(math.Floor(2.8))
    fmt.Println(math.Ceil(2.8))
    fmt.Println(math.Sqrt(16))
    fmt.Println(strutil.Reverse("Hello olleh"))
}
