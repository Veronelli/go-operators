package main

import (
	"fmt"
)

func main() {
	valor1 := 1
	valor2 := 2

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	if valor1 == 1 {
		fmt.Println("Igual")

	} else {
		fmt.Printf("No es igual")
	}

	if (valor1 == 3) && (valor2 == 2) {
		fmt.Println("Es Verdad")
	} else {
		fmt.Println("Es False")
	}

	if (valor1 == 3) || (valor2 == 2) {
		fmt.Println("Es Verdad")
	} else {
		fmt.Println("Es False")
	}

	//Con condiciones
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Par")

	default:
		fmt.Println("Inpar")
	}

	//Con condiciones

	valor := 200
	switch {
	case valor > 100:
		fmt.Println("Mayor a 100")

	case valor < 0:
		fmt.Println("Menos a 0")
	default:
		fmt.Println("Chau")
	}

	//Defer
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			continue
		} else if i == 8 {
			fmt.Println("Es 8")
			break
		}
	}

}
