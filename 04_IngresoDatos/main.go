package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var numero1 int
var numero2 int
var otronumero int
var resultado int
var leyenda string
var err error

func main(){
	fmt.Println("Ingrese numero 1: ")
	//Que tipo de dato recibe y a que variable ira
	//Bug en windows, no toma segundo valor
	//fmt.Scanf("%d", numero1)
	//Solucion bug en windows
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")
	//fmt.Scanf("%d", numero2)
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion de cabecera: ")
	// Creamos un nuevo scanner que se alimentara del "Standard Input", que es el teclado por defecto
	scanner := bufio.NewScanner(os.Stdin)
	//Se evalua si ejecutando el metodo scan del objeto se obtiene algo
	if scanner.Scan() {
		//Que scanner me traiga el texto y lo asigne a la variable
		leyenda = scanner.Text()
	}
	fmt.Println("Tipo de dato de numero1", reflect.TypeOf(numero1))
	
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)


	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese otro numero: ")
	if scanner2.Scan(){
		otronumero, err = strconv.Atoi(scanner2.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto "+ err.Error())
		} 
	}
	fmt.Println("El numero es: ", otronumero)



}