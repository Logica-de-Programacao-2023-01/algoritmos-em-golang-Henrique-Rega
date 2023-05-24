package main

import "fmt"

func main() {
	altura := 0.0
	sexo := ""
	peso := 0.0
	fmt.Println("Insira a altura, sexo, peso REPSECTIVAMENTE:")
	fmt.Scanln(&altura)
	fmt.Scanln(&sexo)
	fmt.Scanln(&peso)
	if peso/(altura*altura) > 25 {
		fmt.Println("Você está acima do peso")
	} else {
		fmt.Println("Você está com o o peso ideal")
	}
}
