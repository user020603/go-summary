package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	Dividend float64 
	Divisor float64 
	Message string
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("error dividing %.2f by %.2f: %s", e.Dividend, e.Divisor, e.Message)
}

func divideCustomError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivisionError{
			Dividend: a,
			Divisor: b,
			Message: "Can not divide by zero",
		}
	}
	return a / b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divideCustomError(5, 0)
	if err != nil {
		fmt.Println("Error Custom:", err)
	} else {
		fmt.Println("Result:", result)
	}
}