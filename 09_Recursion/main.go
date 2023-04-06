package main

import "fmt"

// Recursion => Funcion que se llama a si mismo

func main(){
	exponencial(2)
}

func exponencial ( valor int){
	if valor > 100000000{
		return
	}
	fmt.Println(valor)
	//Recursion, la funcion se llama a si misma
	exponencial(valor*2)
}