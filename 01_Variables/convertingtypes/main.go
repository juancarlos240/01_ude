package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main(){
	
	// Converting Int to Strings using strconv.Itoa
	var numero1 int = 1
	fmt.Println("El tipo de dato inicial de ",numero1, " es : ", reflect.TypeOf(numero1))
	a := strconv.Itoa(numero1)
	fmt.Println("El tipo de dato covertido de ",a, " es :",   reflect.TypeOf(a))
	
	// Converting Strings to Int strconv Atoi 
	texto1 := "1"
	fmt.Println("El tipo de dato inicial de ",texto1, " es : ", reflect.TypeOf(texto1))
	b, err := strconv.Atoi(texto1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El tipo de dato covertido de ",b, " es :",   reflect.TypeOf(b))

	// Converting Strings to Float strconv ParseFloat
	texto2 := "2.2"
	fmt.Println("El tipo de dato inicial de ",texto2, " es : ", reflect.TypeOf(texto2))
	c, err := strconv.ParseFloat(texto2,64)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("El tipo de dato covertido de ",c, " es :",   reflect.TypeOf(c))

	// Converting Strings to Int error

	textoGen := "not a number"
	d, err := strconv.Atoi(textoGen)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(d)


}