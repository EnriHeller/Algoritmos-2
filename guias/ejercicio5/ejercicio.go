package ejercicio5

func (lista *ListaEnlazada[T]) AnteKUltimo(k int) T {

    counter := 0
    var kUltimo T;

    if lista.prim != nil{
        kUltimo = lista.prim.dato
    }

    for lista.prim.prox != nil{

        nodoActual := lista.prim

        lista.prim = nodoActual.prox
        
        if counter == k {
            kUltimo = nodoActual.dato
            counter = 0
        }
        counter ++
        
    }


    return kUltimo
}