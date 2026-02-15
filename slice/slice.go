package main

import "fmt"

func main(){
	var slice [] int
	fmt.Println(slice)

/* 	slice = make([]string, 4)
	fmt.Println(slice) */

	slice = append(slice, 5, 76)
	fmt.Println(slice)
	fmt.Println(len(slice))

	slice = append(slice, 18)
	fmt.Println(slice)
	fmt.Println(len(slice))
	
	//slice := []string{"Angular", "React", "Vue", "Svelte"}
}