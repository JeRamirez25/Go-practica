package main

import "fmt"

func main() {
	result := sum(2, 8)
	fmt.Println(result)

	result2 := sum2(2, 8, 5)
	fmt.Println(result2)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(nums ...int) int {
	var total int
	for _, num := range nums{
		total += num
	}
	return total
}