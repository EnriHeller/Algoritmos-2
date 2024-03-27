package main

import (
	"fmt"
	TDASPila "tdas/pila"
)

func main(){
	pilaPrueba := TDASPila.CrearPilaDinamica[int]()
	pilaPrueba.Apilar(3)
	pilaPrueba.Apilar(2)
	fmt.Println(pilaPrueba.VerTope())
	pilaPrueba.Desapilar()
	fmt.Println(pilaPrueba.VerTope())
	fmt.Println(pilaPrueba.EstaVacia())
}