package calculadora

import "math"

func (calculadora *calculadoraDs) sumar() (int64, bool) {

	elementos, err := calculadora.desapiloN(2)

	return elementos[1] + elementos[0], err
}

func (calculadora *calculadoraDs) restar() (int64, bool) {

	elementos, err := calculadora.desapiloN(2)
	return elementos[1] - elementos[0], err
}

func (calculadora *calculadoraDs) multiplicar() (int64, bool) {

	elementos, err := calculadora.desapiloN(2)
	return elementos[1] * elementos[0], err
}

func (calculadora *calculadoraDs) dividir() (int64, bool) {

	elementos, err := calculadora.desapiloN(2)

	if elementos[0] == 0 {
		return 0, true
	}

	return elementos[1] / elementos[0], err
}

func (calculadora *calculadoraDs) raiz() (int64, bool) {

	elementos, err := calculadora.desapiloN(1)

	if elementos[0] < 0 {
		return 0, true
	}

	resultado := int64(math.Sqrt(float64(elementos[0])))

	return resultado, err
}

func (calculadora *calculadoraDs) potencia() (int64, bool) {

	elementos, err := calculadora.desapiloN(2)

	if elementos[0] < 0 {
		return 0, true
	}

	resultado := int64(math.Pow(
		float64(elementos[1]),
		float64(elementos[0])),
	)

	return resultado, err
}

func (calculadora *calculadoraDs) logaritmo() (int64, bool) {

	elementos, err := calculadora.desapiloN(2)
	a, base := elementos[1], elementos[0]

	if base < 2 {
		return 0, true
	}

	resultado := int64(math.Log(float64(a)) / math.Log(float64(base)))

	return resultado, err
}

func (calculadora *calculadoraDs) ternario() (int64, bool) {

	elementos, err := calculadora.desapiloN(3)

	if elementos[2] == 0 {
		return elementos[0], err
	}
	return elementos[1], err
}
