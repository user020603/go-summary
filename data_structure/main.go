package main

import "fmt"

type Person struct {
	Name string 
	Age int
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1 : 3]
	fmt.Println(s)

	fmt.Println("##################")

	s = make([]int, 0, 1)
	for i := 0; i <= 1300; i++ {
		s = append(s, i)
		fmt.Printf("Len: %d, Cap: %d", len(s), cap(s))
	}

	fmt.Println("##################")

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m["one"])

	fmt.Println("##################")

	p := Person{Name: "Alice", Age: 22}
	p.Age = 25
	fmt.Println(p)

	fmt.Println("##################")

	ch := make(chan int, 1)
	ch <- 42
	fmt.Println(<-ch)
}