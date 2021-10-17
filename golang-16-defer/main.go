package main

import "fmt"

func main() {
	//defer follows the LIFo order: Output is Hello Two One World
	//0,1,2,3,4
	//hello, 4,3,2,1,0,two,one,world
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}
