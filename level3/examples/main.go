package main

import "fmt"

// Person struct to hold information about a person
type Person struct {
	Name string
	Age  int
}

// Greet method for the Person struct
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	// Create an instance of Person
	person := Person{Name: "Alice", Age: 30}

	// Call the Greet method
	person.Greet()
}
