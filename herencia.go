package main

import (
	"fmt"
)

// Declaramos la estructura
type Persona struct {

	// definimos los atributos
	nombre    string
	apellido  string
	edad      int
	carrera   string
	matricula int
}

// Funcion de descripción
func (persona Persona) imprimir_detalles() {

	fmt.Printf("Nombre completo: %s %s", persona.nombre, persona.apellido)
	fmt.Printf("\nCarrera: %s", persona.carrera)
	fmt.Printf("\nEdad: %d \n", persona.edad)
}

type Estudiante struct {
	//Aplicamos herencia de la estructura persona
	Persona
}

type Maestro struct {
	Persona
}

func main() {
	//Instanciamos la estructura
	estudiante1 := Estudiante{Persona{"Alan", "Perez", 22, "Lic. en Derecho", 12345}}

	//Imprimimos la estructura
	estudiante1.imprimir_detalles()
}
