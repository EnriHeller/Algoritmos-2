package ejercicio5

type nodoLista[T any] struct {
    prox *nodoLista[T]
    dato T
}

type ListaEnlazada[T any] struct {
    prim *nodoLista[T]
}

// Función para crear una nueva lista enlazada vacía
func NuevaListaEnlazada[T any]() *ListaEnlazada[T] {
    return &ListaEnlazada[T]{prim: nil}
}

// Función para agregar un elemento al inicio de la lista
func (lista *ListaEnlazada[T]) AgregarInicio(dato T) {
    nuevoNodo := &nodoLista[T]{dato: dato, prox: lista.prim}
    lista.prim = nuevoNodo
}

// Función para agregar un elemento al final de la lista
func (lista *ListaEnlazada[T]) AgregarFinal(dato T) {
    nuevoNodo := &nodoLista[T]{dato: dato, prox: nil}
    if lista.prim == nil {
        lista.prim = nuevoNodo
        return
    }
    actual := lista.prim
    for actual.prox != nil {
        actual = actual.prox
    }
    actual.prox = nuevoNodo
}

// Función para eliminar el primer elemento de la lista
func (lista *ListaEnlazada[T]) EliminarInicio() {
    if lista.prim == nil {
        return
    }
    lista.prim = lista.prim.prox
}

// Función para verificar si la lista está vacía
func (lista *ListaEnlazada[T]) EstaVacia() bool {
    return lista.prim == nil
}

// Función para obtener la longitud de la lista
func (lista *ListaEnlazada[T]) Longitud() int {
    contador := 0
    actual := lista.prim
    for actual != nil {
        contador++
        actual = actual.prox
    }
    return contador
}

// Función para obtener un slice con los elementos de la lista
func (lista *ListaEnlazada[T]) ObtenerElementos() []T {
    elementos := make([]T, 0)
    actual := lista.prim
    for actual != nil {
        elementos = append(elementos, actual.dato)
        actual = actual.prox
    }
    return elementos
}
