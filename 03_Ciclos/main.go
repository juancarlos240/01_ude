package main

import "fmt"

func main(){
	for i := 1; i <=10; i++{
		//fmt.Println(i)
	}
	// output: 1 2 3 4 5 6 7 8 9 10

	numero := 0
	for{
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100{
			break
		}

	}
}