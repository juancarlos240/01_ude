package main

import (
	"fmt"

	"github.com/juancarlos240/Go_01/ejercicios"
	//"github.com/juancarlos240/Go_01/varibles"
)

func main() {
	/*estado, text := varibles.ConviertoaTexto(9889)
	fmt.Println(estado)
	fmt.Println(text)*/

	//os := runtime.GOOS
	/*if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Es Linux")
	} else if os == "windows" {
		fmt.Println("Es windows")
	} else {
		fmt.Println(os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/

	numero, texto := ejercicios.Myfunction("45")
	fmt.Println(numero)
	fmt.Println(texto)

}
