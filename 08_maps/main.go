package main

import "fmt"

func main() {
    // Define map[
    // emails := make(map[string]string)

    // Assign key, values
    // emails["Alex"] = "alex@gmail.com"
    // emails["Tom"] = "tom@gmail.com"
    // emails["Ken"] = "ken@gmail.com"

    // Declare and add key values
    emails := map[string]string{"Alex":"alex@gmail.com", "Ken":"ken@gmail.com"}

    emails["Tom"] = "tom@gmail.com"

    fmt.Println(emails)
    fmt.Println(len(emails))
    fmt.Println(emails["Alex"])

    //Delete from map
    delete(emails, "Ken")
    fmt.Println(emails)
}
