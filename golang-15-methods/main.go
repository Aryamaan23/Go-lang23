package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang...")
	// no inheritance in golang;No super or parent

	aryamaan := User{"Aryamaan23", "aryamaan@go.dev", true, 21}
	fmt.Println(aryamaan)
	fmt.Printf("aryamaan details are: %+v\n", aryamaan)
	fmt.Printf("Name is %v and Email is %v: \n", aryamaan.Name, aryamaan.Email)
	aryamaan.GetStatus()
	aryamaan.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)

}
