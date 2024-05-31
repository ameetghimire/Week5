package main

import "fmt"

func main() {
	name := "Amit"
	message := "Hello"
	age := 28
	greeting(message, name, age)
}
func greeting(message string, name string, age int) {
	//var name string = "Amit"
	fmt.Print(message, name, ".You are", age, "old \n")
}
