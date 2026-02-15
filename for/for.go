package main

import "fmt"

func main(){
	/* EJEMPLO DE BUCLE INFINITO
	i := 1
	for{
		fmt.Println("Loop", i)
		i++
		break
	} */

	/* 
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	} */


	// El orden de ejecucion es el siguiente:
	// 1. Inicializar la variable
	// 2. Verifica la condición
	// 3. Incrementa
	/* for i:=3; i<=10; i++{
		fmt.Println(i)
	} */

	myArray := [4] string {"Angular", "React", "Vue", "Svelte"}

	for key, value := range myArray {
		fmt.Println(key, value)
	}

	for _, value := range myArray {
		fmt.Println(value)
	}

}