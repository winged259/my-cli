package cmd

import (
    "fmt"
    "strconv"
)

func Add(first string, second string) (result string) {
    num1, err := strconv.ParseFloat(first, 64)
    if err != nil {
        fmt.Println("Error: First value invalid")
    }
    num2, err := strconv.ParseFloat(second, 64)
    if err != nil {
        fmt.Println("Error: Second value invalid")
    }
    return fmt.Sprintf("%f", num1+num2)
}

func Subtract(first string, second string) (result string) {
    num1, err := strconv.ParseFloat(first, 64)
    if err != nil {
        fmt.Println("Error: First value invalid")
    }
    num2, err := strconv.ParseFloat(second, 64)
    if err != nil {
        fmt.Println("Error: Second value invalid")
    }
    return fmt.Sprintf("%f", num1-num2)
}


func Multiply(first string, second string, roundUp bool) (result string) {
    num1, err := strconv.ParseFloat(first, 64)
    if err != nil {
        fmt.Println("Error: First value invalid")
    }
    num2, err := strconv.ParseFloat(second, 64)
    if err != nil {
        fmt.Println("Error: Second value invalid")
    }
    if roundUp {
        return fmt.Sprintf("%.2f", num1*num2)
    }
    return fmt.Sprintf("%f", num1*num2)
}
