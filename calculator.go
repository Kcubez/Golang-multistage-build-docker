// calculator.go
package main

import (
    "fmt"
)

func main() {
    var a, b float64
    var op string
    fmt.Println("Enter first number:")
    fmt.Scanln(&a)
    fmt.Println("Enter second number:")
    fmt.Scanln(&b)
    fmt.Println("Enter operation (+, -, *, /):")
    fmt.Scanln(&op)

    switch op {
    case "+":
        fmt.Printf("Result: %f\n", a+b)
    case "-":
        fmt.Printf("Result: %f\n", a-b)
    case "*":
        fmt.Printf("Result: %f\n", a*b)
    case "/":
        fmt.Printf("Result: %f\n", a/b)
    default:
        fmt.Println("Invalid operation")
    }
}
