package main

import "fmt"

func main() {
	result, ok := divide(100, 10)
	fmt.Println(result, ok)
}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}