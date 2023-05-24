package main

import "fmt"

func main() {
	fmt.Println("Insira o número desejado:")
	x := 0
	fmt.Scan(&x)
	if x % 2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é impar")
	}
}
