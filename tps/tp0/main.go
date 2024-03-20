package main

import (
	"fmt"
	"tp0/ejercicios"
)
func main()  {
	/*
	
	//1
	x, y := 1, 2
	ejercicios.Swap(&x, &y)

	fmt.Println(x,y) 

	//2
	arreglo := [4]int{1,12,3,5}
	resultado := ejercicios.Maximo(arreglo[:])
	fmt.Println(resultado)
	
	//3
	arreglo1 := [4]int{1,34,23,4}
	arreglo2 := [3]int{1,34,8}

	resultado := ejercicios.Comparar(arreglo1[:], arreglo2[:])
	fmt.Println(resultado)

	//4
	arreglo4 := [5]int{32,5,1,8,6}
	ejercicios.Seleccion(arreglo4[:])
	fmt.Println(arreglo4)

	//5
	array := [4]int{1,2,5,8}
	resultado := ejercicios.Suma(array[:])
	fmt.Println(resultado)

	//6
	cadena := "oso"
	resultado := ejercicios.EsCadenaCapicua(cadena)
	fmt.Println(resultado)
	*/

	//2
	arreglo := []int{-2000, -1500, -1000, -3000}
	resultado := ejercicios.Maximo(arreglo[:])
	fmt.Println(resultado)
}



