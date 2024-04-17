package calculadora

import (
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

type calculadoraDs struct{
	operandos TDAPila.Pila[int64]
}

func CrearCalculadoraDs() calculadoraDs{
	return calculadoraDs{
		operandos:TDAPila.CrearPilaDinamica[int64](),
	}
}

func (calculadora *calculadoraDs) Operar(comando string) (int64, bool){
	elementos := strings.Fields(comando)

	for _, elem := range elementos {
		operando, err := strconv.Atoi(elem)
		esNumero := err == nil

		if esNumero{
			calculadora.operandos.Apilar(int64(operando))
		}else{
			resultado, err := calculadora.aplicarOperacion(elem)

			if err{
				return 0, true
			}
			calculadora.operandos.Apilar(resultado)
		}
	}

	return calculadora.operandos.VerTope(), false
}

func (calculadora *calculadoraDs) aplicarOperacion(operador string) (int64,bool){

	switch operador {
	case "+":
		return calculadora.sumar()
	case "-":
		return calculadora.restar()
	case "*":
		return calculadora.multiplicar()
	case "/":
		return calculadora.dividir()
	case "sqrt":
		return calculadora.raiz()
	case "^":
		return calculadora.potencia()
	case "log":
		return calculadora.logaritmo()
	case "?":
		return calculadora.ternario()
	}

	return 0, true
}



