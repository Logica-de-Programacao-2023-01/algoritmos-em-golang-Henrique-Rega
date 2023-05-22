package main

import "fmt"

func main() {
	x := 0
	y := 0
	fmt.Scan(&x, &y)
	if x > y {
		fmt.Print("X é maior que Y")
	} else {
		fmt.Print("Y é maior que X")
	}

}
