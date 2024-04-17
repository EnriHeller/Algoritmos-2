package main

import (
	"bufio"
	"fmt"
	"os"
	"tp1/calculadora"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		calculadoraDs := calculadora.CrearCalculadoraDs()
		comando := entrada.Text()
		resultado, err := calculadoraDs.Operar(comando)
		if err {
			fmt.Println("ERROR")
		} else {
			fmt.Println(resultado)
		}
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Printf("Error al leer entrada: %s", errEntrada)
	}
}
