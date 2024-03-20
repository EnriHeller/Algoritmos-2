package ejercicios

import "strings"

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {

	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {

	if len(vector) == 0 {
		return -1
	}

	indiceActual := 0

	for indice, valor := range vector {
		if valor > vector[indiceActual] {
			indiceActual = indice
		}
	}

	return indiceActual
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.

func compararInterno(vector1 []int, vector2 []int, lenMenor int) int {

	for i := range lenMenor {
		if vector1[i] > vector2[i] {
			return 1
		}

		if vector2[i] > vector1[i] {
			return -1
		}
	}

	return 0
}

func Comparar(vector1 []int, vector2 []int) int {

	if len(vector1) < len(vector2) {
		rParcial := compararInterno(vector1, vector2, len(vector1))
		if rParcial == 0 {
			return -1
		}
	} else if len(vector2) < len(vector1) {
		rParcial := compararInterno(vector1, vector2, len(vector2))
		if rParcial == 0 {
			return 1
		}
	}

	return compararInterno(vector1, vector2, len(vector1))
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {

	iUltimo := len(vector) - 1

	for iUltimo > 0 {
		iMax := Maximo(vector[:iUltimo+1])
		Swap(&vector[iMax], &vector[iUltimo])
		iUltimo = iUltimo - 1
	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func SumarRestantes(vector []int, total int) int {

	if len(vector) == 0 {
		return total
	}

	n := len(vector) - 1
	total += vector[n]

	return SumarRestantes(vector[:n], total)
}

func Suma(vector []int) int {
	return SumarRestantes(vector, 0)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func invertirCadena(cadena []string, cadenaInvertida string) string {

	if len(cadena) == 0 {
		return cadenaInvertida
	}

	n := len(cadena) - 1
	ultimoC := cadena[n]
	cadenaInvertida += ultimoC

	return invertirCadena(cadena[:n], cadenaInvertida)
}

func EsCadenaCapicua(cadena string) bool {

	caracteres := strings.Split(cadena, "")
	cadenaInvertida := invertirCadena(caracteres, "")

	return cadena == cadenaInvertida
}
