package main

func main(){

}

// Nombre de la funcion "uno", recibe un parametro y retorna un valor 
func uno(numero int) (z int) {
	z = numero * 2
	return
}

//Funcion que devuelve 2 datos
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}


//Parametros variables, va a recibir parametros enteros, pero no sabemos cuantos
func Calculo(numero ...int) int {
	total := 0
	//Iterar por cada parametro recibido
	//El guion bajo aloja variables que no voy a usar, la devuelve range, y no la voy a usar
	for _, num := range numero {
		total = total + num
	}

	return total
}


/*
- En go no existe la sobrecarga de metodos ni de funciones, si tiene implementado el numero variable de parametros
- Range devuelve 2 valores, se usa mucho en vectores, para recorrer todos los elementos de una lista de datos. El primer valor, es el numero de elemento actual y
el segundo es el elemento en si
*/