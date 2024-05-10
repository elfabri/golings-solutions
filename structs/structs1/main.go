// structs1
// Make me compile!

package main

import "fmt"

type Person struct {
    name string
    age int32
}

var person Person = Person {name: "Fabricio", age: 25}

func main() {
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
