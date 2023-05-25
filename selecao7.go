package main

import "fmt"

func main() {
	salário := 0.0
	fmt.Print("Digite o salário: ")
	fmt.Scanln(&salário)
	if salário < 1000.00 {
		fmt.Printf("o novo salário é de: R$%.2f\n ", salário*1.10)
	} else {
		fmt.Printf("o novo salário é de: R$%.2f\n", salário*1.05)
	}
}
