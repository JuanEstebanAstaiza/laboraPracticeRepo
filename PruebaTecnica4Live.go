package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Función para generar una matriz de tamaño aleatorio nxm

func generarMatriz(filas, columnas int) [][]int {
	matriz := make([][]int, filas)
	for i := 0; i < filas; i++ {
		matriz[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}

// Función de imprimir la matriz

func imprimirMatriz(matriz [][]int) {
	for _, fila := range matriz {
		fmt.Println(fila)
	}
}

//Funcion de multiplicación concurrente

func multiplicarMatricesConcurrente(matrizA, matrizB [][]int, resultado chan<- [][]int) {
	filasA := len(matrizA)
	columnasA := len(matrizA[0])
	columnasB := len(matrizB[0])

	resultadoMatriz := make([][]int, filasA)
	for i := 0; i < filasA; i++ {
		resultadoMatriz[i] = make([]int, columnasB)
	}

	for i := 0; i < filasA; i++ {
		for j := 0; j < columnasB; j++ {
			go func(i, j int) {
				suma := 0
				for k := 0; k < columnasA; k++ {
					suma += matrizA[i][k] * matrizB[k][j]
				}
				resultadoMatriz[i][j] = suma
			}(i, j)
		}
	}
	//funcion de espera de Goroutines
	for x := 0; x < filasA; x++ {
		for z := 0; z < columnasB; z++ {
			<-time.After(10 * time.Millisecond)
		}
	}

	resultado <- resultadoMatriz
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//definir tamaño de matriz

	filasA := 3
	columnasA := 2
	filasB := 2
	columnasB := 3

	//generar matrices

	matrizA := generarMatriz(filasA, columnasA)
	matrizB := generarMatriz(filasB, columnasB)

	//imprimir las matrices iniciales

	fmt.Println("Matriz A: ")
	imprimirMatriz(matrizA)
	fmt.Println("Matriz B: ")
	imprimirMatriz(matrizB)

	// Crear canal resultado

	resultado := make(chan [][]int)

	//multiplicar matrices concurrentes

	go multiplicarMatricesConcurrente(matrizA, matrizB, resultado)

	//Espera e impresion del resultado

	fmt.Println("Resultado del producto cruz: ")
	fmt.Println(<-resultado)
}
