package main

import (
	"fmt"
	"strings"
)

// Struct del ahorcado

type JuegoAhorcado struct {
	palabra       string   // Palabra a adivinar
	letras        []string // Letras adivinadas correctamente
	intentos      int      // número de intentos restantes
	intentosMax   int      // Número maximo de intentos
	palabraOculta string   //Represantacion de la palabra oculta con guiones

}

//crear funcion para el juego del ahorcado

func NuevoJuegoAhorcado(palabra string, intentosMax int) *JuegoAhorcado {

	palabra = strings.ToUpper(palabra) // Convertir las palabras en mayusculas permitiendo que las validaciones sean insensibles a letras minusculas
	palabraOculta := strings.Repeat("_ ", len(palabra))
	return &JuegoAhorcado{

		palabra:       palabra,
		letras:        make([]string, 0),
		intentos:      intentosMax,
		intentosMax:   intentosMax,
		palabraOculta: palabraOculta,
	}
}

//funcion para adivinar las letras

func (j *JuegoAhorcado) AdivinarLetra(letra string) {

	letra = strings.ToUpper(letra) //Convertimos a mayuscula
	if !strings.Contains(j.palabra, letra) {
		j.intentos--
	} else {
		j.letras = append(j.letras, letra)
		//Actualizacion de la represantacion de la palabra oculta
		palabraOculta := ""
		for _, char := range j.palabra {
			if contains(j.letras, string(char)) { //Crear funcion contains
				palabraOculta += string(char) + " "
			} else {
				palabraOculta += "_ "
			}
		}

		j.palabraOculta = palabraOculta
	}
}

//Funcion de verificacion de victoria

func (j *JuegoAhorcado) HaGanado() bool {
	return j.palabraOculta == j.palabra
}

//Funcion de verificacion de derrota

func (j *JuegoAhorcado) HaPerdido() bool {
	return j.intentos <= 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	//Crear un nuevo juego con la palabra "GOL"

	juego := NuevoJuegoAhorcado("GOL", 5)

	//Ciclo de juego

	for {
		fmt.Println("Intentos restantes: ", juego.intentos)
		fmt.Println("Palabra: ", juego.palabraOculta)

		//solicitud de jugador para adivinar una letra

		var letra string
		fmt.Print("Adivina una letra: ")
		fmt.Scanln(&letra)

		//Adivinar la letra

		juego.AdivinarLetra(letra)

		//verificar si se ha ganado o perdido

		if juego.HaGanado() {
			fmt.Println("¡Has ganado! La palabra era: ", juego.palabra)
			break
		} else if juego.HaPerdido() {
			fmt.Println("¡Has perdido! La palabra era: ", juego.palabra)
			break
		} else if juego.HaGanado() {
			fmt.Println("¡Has ganado! La palabra era: ", juego.palabra)
			break
		}
	}
}
