package ds

import (
	"fmt"
	"math"
	TDACola "tdas/cola"
	"strconv"
	"strings"
)

func Sumar(a int64, b int64) int64 {
	return a + b
}

func Restar(a int64, b int64) int64 {
	return a - b
}

func Multiplicar(a int64, b int64) int64 {
	return a * b
}

func Dividir(a int64, b int64) int64 {
	return a / b
}

func Raiz(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

func Potenciar(a int64, b int64) int64 {
	return int64(math.Pow(
		float64(a),
		float64(b)),
	)
}

func Log(numero int64, base int64) int64 {

	resultado := math.Log(float64(numero)) / math.Log(float64(base))
	return int64(resultado)
}

func Ternario(a int64, b int64, c int64) int64 {

	if a == 0 {
		return c
	}
	return b
}

func obtenerOps(operacion string) (TDACola.Cola[int64], TDACola.Cola[string]) {
	operandos := TDACola.CrearColaEnlazada[int64]()
	operadores := TDACola.CrearColaEnlazada[string]()

	elementos := strings.Split(operacion, " ")

	for _, elem := range elementos {

		if strings.Contains("+ - * / ^ sqrt log ?",elem){

			operadores.Encolar(elem)

		}else{
			numero, err := strconv.Atoi(elem)

			if err != nil {
				fmt.Println("Error al convertir el caracter a número:", err)
				continue
			}
			operandos.Encolar(int64(numero))
		}
	}

	return operandos, operadores
}

func aplicarOperacion(operacion func(a,b int64) int64, operandos TDACola.Cola[int64]) int64{
	
	if !operandos.EstaVacia(){
		return operacion(operandos.VerPrimero(),operandos.Desencolar())
	}else{
		fmt.Println("ERROR")
		return 0
	}
}

func Operar(operacion string) int64 {

	operandos, operadores := obtenerOps(operacion)
	var resultado int64
	
	if !operandos.EstaVacia(){
		resultado = operandos.Desencolar()
	}


	for !operadores.EstaVacia(){

		operador := operadores.Desencolar()
		switch operador {
		case "+":
			aplicarOperacion(Sumar,operandos)
			//resultado = Sumar(resultado, operandos.Desencolar())
		case "-":
			resultado = Restar(resultado, operandos.Desencolar())
		case "*":
			resultado = Multiplicar(resultado, operandos.Desencolar())
		case "/":
			resultado = Dividir(resultado, operandos.Desencolar())
		case "^":
			resultado = Potenciar(resultado, operandos.Desencolar())
		case "log":
			resultado = Log(resultado, operandos.Desencolar())
		case "sqrt":
			resultado = Raiz(resultado)
		case "?":
			resultado = Ternario(resultado, operandos.Desencolar(), operandos.Desencolar())
		}
	}

	return resultado

}
