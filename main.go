package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 1. Leer el input (GitHub antepone "INPUT_" al nombre de la variable en mayúsculas)
	name := os.Getenv("INPUT_WHO_TO_GREET")
	if name == "" {
		name = "Mundo" // Valor por defecto si no se proporciona
	}

	// 2. Imprimir en los logs de la Action
	fmt.Printf("¡Hola, %s!\n", name)

	// 3. Escribir el output para que otros pasos lo puedan usar
	outputFilePath := os.Getenv("GITHUB_OUTPUT")
	if outputFilePath != "" {
		file, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Error al abrir el archivo GITHUB_OUTPUT: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		currentTime := time.Now().Format(time.RFC3339)
		// El formato para escribir outputs es "llave=valor\n"
		outputString := fmt.Sprintf("time=%s\n", currentTime)
		
		if _, err := file.WriteString(outputString); err != nil {
			fmt.Printf("Error al escribir el output: %v\n", err)
			os.Exit(1)
		}
	}
}