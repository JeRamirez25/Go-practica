package main

import "fmt"

// Un Struct es parecido a las clases de python

type Profession string

// Declarar constantes con valores predeterminados

const (
	Lawyer Profession = "Lawyer"
	Doctor Profession = "Doctor"
	Programmer Profession = "Programmer"
)

type Person struct {
	Name string
	Age  int
	Job Profession
}

func main() {
	var p Person = Person{Name: "Juan", Age: 21, Job: Programmer}

	q := Person{Name: "Esteban", Age: 21, Job: Doctor} // Se puede iniciarlizar de esta manera de una manera mas directa
	
	/* p.Name = "Juan"
	p.Age = 21 */

	fmt.Println(p)
	fmt.Println(q)
}
