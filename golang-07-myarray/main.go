package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of Fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list is ", vegList)
	fmt.Println("Length of Veggie list is ", len(vegList))

}
