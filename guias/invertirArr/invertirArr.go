package invertirArr

import("tdas/pila")

func InvertirArreglo[T any](arr []T) {
	pila := pila.CrearPilaDinamica[T]()

	for i:= 0; i<len(arr); i++{
		pila.Apilar(arr[i])
	}

	for i:= 0; i<len(arr); i++{
		arr[i] = pila.Desapilar()
	}
}