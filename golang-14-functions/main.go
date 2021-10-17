package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang!")
	greeter()

	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro Result is: ", proResult)
	fmt.Println("Pro Message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro Result function"
}

func greeterTwo() {
	fmt.Println("Another method of Greeting")
}

func greeter() {
	fmt.Println("Namaste from Golang")
}
