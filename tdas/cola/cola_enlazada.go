package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodoCola[T any]() *nodoCola[T] {

	return new(nodoCola[T])
}

func CrearColaEnlazada[T any]() Cola[T] {

	nuevaCola := new(colaEnlazada[T])
	return nuevaCola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {

	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {

	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {

	nuevoNodo := crearNodoCola[T]()
	nuevoNodo.dato = elem

	if cola.EstaVacia() {
		cola.primero = nuevoNodo
		cola.ultimo = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
		cola.ultimo = nuevoNodo
	}
}

func (cola *colaEnlazada[T]) Desencolar() T {

	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	primerElemento := cola.primero.dato
	cola.primero = cola.primero.prox

	if cola.primero == nil {
		cola.ultimo = nil
	}

	return primerElemento
}
