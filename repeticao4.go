package main

import "fmt"

func main() {
	for i := 0; i < 33; i++ {
		for j := i; j%3 == 0; j++ {
			fmt.Println(j)
		}
	}
}
