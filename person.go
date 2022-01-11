package main

import "fmt"

type Person struct {
	name string
	age  int
}

// CheckAge function is return true and person details if person age > 22 and return false and empty if age is < 22
func CheckAge(p Person) (res bool, val Person) {
	if p.age < 22 {
		return
	}
	return true, p
}
func main() {
	v := Person{name: "Ram", age: 21}
	fmt.Println(CheckAge(v))

}
