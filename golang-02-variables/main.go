package main

import "fmt"

// not allowed to use walrus operator outside method jwtToken :=300000

// If first alphabet of the variable is capital then it is a public keyword
const LoginToken string = "ghdfosfofsdof"

func main() {
	var username string = "Aryamaan23"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.455423232323232323
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "aryamaan23.com"
	fmt.Println(website)

	//no var style
	numberofUser := 300000
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
