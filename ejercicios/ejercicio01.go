package ejercicios

import (
	"fmt"
	"strconv"
)

func Myfunction(valor string) (int, string) {
	value, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println(err.Error())
	}
	if value > 100 {
		return value, "Es mayor a 100"
	} else {
		return value, "Es menor a 100"
	}

}
