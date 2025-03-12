// package
package main

import "fmt"

// Variable
var userName string

// Unexported Function
func calculateSum(a, b int) int {
	return a + b
}

// Exported Function
func PrintMessage(msg string) {
	fmt.Println(msg)
}

// Struct
type User struct {
	Name  string
	Email string
	Age   int
}

// Interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Constant
const MAX_BUFFER_SIZE = 1024
const PI = 3.14

func main() {
	userName = "Thanh"
	fmt.Println(userName)
	fmt.Println(calculateSum(1, 2))
}
