package main

import (
	"fmt"
	"tp1/calculadora"
	"os"
	"bufio"
)

//./ds < oper.txt > dc.txt

/*func main(){
	operacion := "2 2 + +"
	Calculadora := calculadora.CrearCalculadoraDs()
	resultado, err := Calculadora.Operar(operacion)
	if(err){
		fmt.Println("llego el error al final", err)
	}else{
		fmt.Println("llego resultado al final",resultado)
	}
}*/


func main() {
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		calculadoraDs := calculadora.CrearCalculadoraDs()
		comando := entrada.Text()
		resultado, err := calculadoraDs.Operar(comando)
		if err{
			fmt.Println("ERROR")
		}else{
			fmt.Println(resultado)
		}
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Printf("Error al leer entrada: %s", errEntrada)
	}
}


