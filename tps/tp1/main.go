package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		result, err := processLine(line)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

func processLine(line string) (string, error) {
	// Aquí procesarías la línea y devolverías el resultado y un error si corresponde
	// Por ahora, simplemente devolvemos la misma línea como resultado
	return line, nil
}
