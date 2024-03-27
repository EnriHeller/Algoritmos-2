package pila

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const capInicial int = 6

func CrearPilaDinamica[T any]() Pila[T] {

	nuevaPila := new(pilaDinamica[T])
	nuevaPila.datos = make([]T, capInicial)
	return nuevaPila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {

	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {

	iTope := pila.cantidad - 1
	return pila.datos[iTope]
}

func (pila *pilaDinamica[T]) redimensionar(tam int) {
	nuevoArreglo := make([]T, tam)
	copy(nuevoArreglo, pila.datos)
	pila.datos = nuevoArreglo
}

func (pila *pilaDinamica[T]) checkRedimension() {

	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) * 2)
	}

	if pila.cantidad*4 <= cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) / 2)
	}

}

func (pila *pilaDinamica[T]) Apilar(elem T) {

	pila.datos[pila.cantidad] = elem
	pila.cantidad++
	pila.checkRedimension()
}

func (pila *pilaDinamica[T]) Desapilar() T {

	ultimoElemento := pila.VerTope()
	pila.cantidad--
	pila.checkRedimension()
	return ultimoElemento
}
