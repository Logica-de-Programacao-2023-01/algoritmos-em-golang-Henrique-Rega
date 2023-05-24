package main

import "fmt"

func main() {
	var x int
	fmt.Println("insira o número")
	fmt.Scanln(&x)
	if x%3 == 0 {
		if x%5 == 0 {
			fmt.Println("O número é divisível por 3 e 5 ao mesmo tempo")
		}
	} else {
		fmt.Println("o número não é divisível por 3 e 5 ao mesmo tempo")
	}
}
