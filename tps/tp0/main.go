package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func obtenerArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta, err)
		return nil
	}

	return archivo
}

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

func escribirArchivo(ruta string, arreglo []int) {
	archivo, err := os.Create(ruta)
	printError(err)
	defer archivo.Close()

	datawriter := bufio.NewWriter(archivo)

	for _, numero := range arreglo {
		_, err = datawriter.WriteString(fmt.Sprintf("%d", numero) + "\n")
		if err != nil {
			fmt.Printf("No se pudo guardar el número %d, error: %v", numero, err)
		}
	}
	datawriter.Flush()
}

func imprimirArchivo(ruta string) {
	archivo, err := os.Open(ruta)
	printError(err)
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
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
		escribirArchivo(ruta1, arreglo1)
		imprimirArchivo(ruta1)
	} else if comparacion == -1 {
		ejercicios.Seleccion(arreglo2)
		escribirArchivo(ruta2, arreglo2)
		imprimirArchivo(ruta2)
	}
}
