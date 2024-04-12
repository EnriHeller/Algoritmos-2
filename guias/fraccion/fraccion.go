package fraccion

import(
	"fmt"
)

type Fraccion struct {
	numerador int
	denominador int
}

// CrearFraccion crea una fraccion con el numerador y denominador indicados. 
// Si el denominador es 0, entra en panico.
func CrearFraccion(numerador, denominador int) Fraccion {

	nuevaFraccion := Fraccion{
		numerador:   numerador,
		denominador: denominador,
	}

	if(nuevaFraccion.denominador == 0){
		panic("No se puede dividir por 0")
	}
	
	return nuevaFraccion
}

// Sumar crea una nueva fraccion, con el resultante de hacer la suma de las fracciones originales
func (f Fraccion) Sumar(otra Fraccion) Fraccion {
	nuevoDenominador := f.denominador * otra.denominador

	nuevoNumerador := 
		f.numerador * otra.denominador +
		otra.numerador * f.denominador
	

		nuevaFraccion := Fraccion{
			numerador:   nuevoNumerador,
			denominador: nuevoDenominador,
		}

	return nuevaFraccion
}

// Multiplicar crea una nueva fraccion con el resultante de multiplicar ambas fracciones originales
func (f Fraccion) Multiplicar(otra Fraccion) Fraccion {
	nuevoDenominador := f.denominador * otra.denominador
	
	nuevoNumerador := f.numerador * otra.numerador

	nuevaFraccion := Fraccion{
		numerador:   nuevoNumerador,
		denominador: nuevoDenominador,
	}
	
	return nuevaFraccion
}

// ParteEntera devuelve la parte entera del numero representado por la fracción. 
// Por ejemplo, para "7/2" = 3.5 debe devolver 3. 
func (f Fraccion) ParteEntera() int {
	division := f.numerador/f.denominador
	
	return int(division)
}


// Representacion devuelve una representación en cadena de la fraccion simplificada (por ejemplo, no puede devolverse
// "10/8" sino que debe ser "5/4"). Considerar que si se trata de un número entero, debe mostrarse como tal.
// Considerar tambien el caso que se trate de un número negativo. 

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func (f Fraccion) Representacion() string {
	MCD := gcd(f.numerador, f.denominador)

	nSimplificado := f.numerador / MCD
	dSimplificado := f.denominador / MCD 


	if f.numerador % f.denominador == 0 {
		return fmt.Sprintf("%d",
		f.numerador/f.denominador)
	}

	return fmt.Sprintf("%d/%d", nSimplificado, dSimplificado)

}