package main

import "fmt"

func main() {
	i := 5

	if i < 5 {
		fmt.Println("i is less than 5")
	} else if i == 5 {
		fmt.Println("I is equal to 5")
	} else {
		fmt.Println("I is at least 5")
	}

	fmt.Println("Finish!!")
}