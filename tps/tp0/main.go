package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

//printError imprime los errores generados, exceptuando los de valor "nil".

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//obtenerArchivo devuelve el archivo cuya ruta se indica por parámetro.
//En caso de haber un error, se imprime la ruta del archivo y el código de error asociado.

func obtenerArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta, err)
		return nil
	}

	return archivo
}

//obtenerArreglo devuelve arreglo de números ubicados separadamente línea por línea
//en el archivo pasado por parámetro.

func obtenerArreglo(archivo *os.File) []int {
	var arreglo []int
	s := bufio.NewScanner(archivo)

	for s.Scan() {
		valor, err := strconv.Atoi(s.Text())
		printError(err)

		arreglo = append(arreglo, valor)
	}

	err := s.Err()
	printError(err)

	return arreglo
}

//imprimirArreglo imprime cada número alojado en el arreglo pasado por parámetro, por separado.

func imprimirArreglo(arreglo []int) {
	for _, valor := range arreglo {
		fmt.Println(valor)
	}
}

func main() {
	ruta1, ruta2 := "./archivo1.in", "./archivo2.in"
	archivo1, archivo2 := obtenerArchivo(ruta1), obtenerArchivo(ruta2)

	defer archivo1.Close()
	defer archivo2.Close()

	arreglo1 := obtenerArreglo(archivo1)
	arreglo2 := obtenerArreglo(archivo2)

	comparacion := ejercicios.Comparar(arreglo1, arreglo2)

	if comparacion == 1 {
		ejercicios.Seleccion(arreglo1)
		imprimirArreglo(arreglo1)
	} else if comparacion == -1 {
		ejercicios.Seleccion(arreglo2)
		imprimirArreglo(arreglo2)
	}
}
