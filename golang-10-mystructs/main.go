package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang...")
	// no inheritance in golang;No super or parent

	aryamaan := User{"Aryamaan23", "aryamaan@go.dev", true, 21}
	fmt.Println(aryamaan)
	fmt.Printf("aryamaan details are: %+v\n", aryamaan)
	fmt.Printf("Name is %v and Email is %v: ", aryamaan.Name, aryamaan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
