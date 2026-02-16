package main

import "fmt"

// Se debe declarar desde el inicio el tipo de datos que va a contener y la cantidad de estos datos

func main(){
	var myArray [4] string
	fmt.Println(myArray)

	//Agregar datos al array vacio
	myArray = [4] string{"Angular", "React", "Vue", "Svelte"}
	fmt.Println(myArray)

	// Dar posición
	fmt.Println(myArray[1])

	// Copiar todo el array
	myArray2 := myArray

	// Sobreescribir la posición
	myArray[2] = "VueJS"
	fmt.Println(myArray)

	fmt.Println(myArray2)

	fmt.Println(len(myArray))
}