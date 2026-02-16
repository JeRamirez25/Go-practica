package main

import (
	"fmt"
	"strconv"
)

func main(){
	var myInt int = 42
	var myFloat float64 = float64(myInt)

	fmt.Println(myFloat)

	var myStr string = "123.456"
	// Se usa _ cuando se quiere evitar el nill (null)
	var myFloat2, _ = strconv.ParseFloat(myStr, 64)

	fmt.Println(myFloat2)
}