package main

import (
	
	"fmt"
	"tp1/ds"
	"os"
	"bufio"
)

//./main < oper.txt > dc.txt

/*func main(){
	operacion := "4 sqrt"
	resultado := ds.Operar(operacion)
	fmt.Println(resultado)
}*/


func main() {
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		operacion := entrada.Text()
		resultado := ds.Operar(operacion)
		fmt.Println(resultado)
	}

	if err := entrada.Err(); err != nil {
		fmt.Printf("Error al leer entrada: %s", err)
	}
}


