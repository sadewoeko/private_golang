package main

import "fmt"

func main() {
    fmt.Print("Enter number: ")
    var input string
    fmt.Scanln(&input)
    fmt.Print(input)
}