package main

import "fmt"

func main() {
	var numero int
	fmt.Println("Insira o número: ")
	fmt.Scan(&numero)
	for i := 0; i < 11; i++ {
		fmt.Println(numero, " * ", i, "=", numero*i)
	}
}
