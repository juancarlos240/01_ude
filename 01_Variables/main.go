package main

import (
	"fmt"
	"strconv"
)

// Todo lo declarado aqui, seran variables a nivel del package
var numero int
var texto string

// Los booleanos se inicializan en false
var status bool = false

// Si la variable empieza con Mayuscula, podria ser exportada en paquete
var NumeroPack int = 8

// i = 8 int y j =42 int , toma los tipos explicitomente
var i, j = 8, 42

// Constantes
const pi = 3.1416

func main() {

	// Una variable en Go se puede declarar dos maneras
	//Declarativa
	var intcomun int
	//Asignacion
	intde32 := int32(10)
	fmt.Println(intcomun)
	fmt.Println(intde32)

	// Los numeros se inicializan en cero
	var numero2, numero4, numero5 int

	// Status aqui sera true
	status = true

	// Asignamos directamente y el tipo de variable sera el valor que le di, esta asignacion solo se puede hacer 1 vez por variable, si quiero asignarle un nuevo valor a la variable
	// se usa "=" sin los dos puntos, puedo asignar valores a varias variables a la vez
	numero2x, numero4x, numero5x := 2, 4, 5
	numero2x = 22

	// Otra forma de asignacion del valor a una variable
	var valorInt64 int64 = 0

	//No se puede
	//var numero int = valorInt64
	//Si se puede
	var ejemplo int = int(valorInt64)

	fmt.Println(ejemplo)
	fmt.Println(numero2, numero4, numero5)
	fmt.Println(numero2x, numero4x, numero5x)

	//Imprimir variable int como decimal con fmt
	fmt.Println(fmt.Sprintf("%d", numero2))

	//Convertir entero a String, solo conviente int, si es int64 da error
	fmt.Println(strconv.Itoa(numero2))

	fmt.Println(status)
	mostrarStatus()
	MostrarTexto()

}

// Una funcion o variable pueden ser publicas indicandola primera letra mayuscula
func mostrarStatus() {
	fmt.Println("Hola! Soy la funcion")
	//status aca sera false
	fmt.Println(status)
}

// Lo mismo con las funciones, si empiezan en mayuscula, pueden ser exportados
func MostrarTexto() {
	fmt.Println("Hola! Puedo ser exportado")
	fmt.Println(NumeroPack)
}
