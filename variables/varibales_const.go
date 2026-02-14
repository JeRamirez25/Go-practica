package main

import "fmt"

func main(){
	// Se puede declarar y despues darle valor
	var a int
	a = 10

	// Se puede declarar y darle valor
	var b int = 20

	// Se puede declarar y darle valor (al no tener un tipo definido se asume)
	var c = 30

	//Definir e inicializar, solo dentro de funciones
	d := 40

	fmt.Println(a+b+c)
	fmt.Println(d)

	//Declarar e inicializar dos variables al tiempo -> POCO RECOMENDADO
	var e,f int = 1,2

	fmt.Println(e+f)

	//Declarar constantes
	const name string = "Juan"
	fmt.Println(name)

	// Declarar multiples constantes, si una varible se declara sin valor, por defecto se toma el valor anterior (NO SE PUEDE DAR VALOR DESPUES)
	const (
		g = 35
		h = 40
		i
	)

	fmt.Println(g, h, i)

	const PI = 3.14
	const radius = 5.0
	const area = PI * radius * radius

	fmt.Println("Radio =", area)
}