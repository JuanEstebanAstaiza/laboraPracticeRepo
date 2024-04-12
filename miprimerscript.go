package main

import (
	"fmt"
)

func sumarnumeros() {
	var numero_uno int = 1
	var numero_dos int = 30

	var resultado int = numero_uno + numero_dos

	fmt.Println(resultado)
}

func substractNumbers() {
	var number1 int = 10
	var number2 int = 30
	var result int = number2 - number1
	fmt.Println(result)
}

func MultiplyNumers() {

	var number1 int = 10
	var number2 int = 30
	var result int = number2 * number1

	fmt.Println(result)
}

func divideNumbers() {

	var number1 int = 30
	var number2 int = 10

	var result = number1 / number2

	fmt.Println(result)
}

func DefinirNumerosPares(numero int) {

	if numero%2 == 0 {

		fmt.Println("es un numero entero")
	} else {
		fmt.Println("no es un numero entero19")
	}
}

func VocalOrConsonant(char string) string {
	switch char {
	case "a", "e", "i", "o", "u":
		return "vocal"
	default:
		return "consonant"
	}
}

func CalcularAreaCirculo() {

}
