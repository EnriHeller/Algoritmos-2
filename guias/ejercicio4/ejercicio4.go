package ejercicio4

type composicionFunc struct {
	funciones []func(float64) float64;
}

//O(1)
/*func CrearComposicion() ComposicionFunciones {
    nuevaComposicion := new(composicionFunc)
	nuevaComposicion.funciones = make([]func(float64) float64, 0)
	return nuevaComposicion
} */

// Pritimitivas

//O(1)
func (comp *composicionFunc) AgregarFuncion(funcion func(float64) float64){
		comp.funciones= append(comp.funciones, funcion)
}



//O(n)
func (comp *composicionFunc) Aplicar(valor float64) float64{

	resultado := valor

	for i:= len(comp.funciones); i<0; i--{
		funcionActual := comp.funciones[i]
		funcionActual(resultado)
	}

	return resultado
} 