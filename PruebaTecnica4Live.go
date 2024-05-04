package main

/*import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)*/

/*import (
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
}*/

//primero creemos los imports que necesitamos

/*import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//creamos la estructura del objeto competidor

type Competidor struct {
	nombre    string
	velocidad int
}

// Vamos a crear la funcion de simulacion de progreso de los competidores en la carrera

func correrCompetidor(competidor Competidor, pista chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for distancia := 0; distancia < 100; {
		avance := rand.Intn(competidor.velocidad)
		pista <- avance
		distancia += avance
		time.Sleep(100 * time.Millisecond) //Aqui vamos a simular el tiempo de la carrera

	}

	fmt.Printf("%s ha terminado la carrera! \n", competidor.nombre)
}

//El time sleep simula la duracion de la carrera

func main() {
	rand.Seed(time.Now().UnixNano())
	//es simplemente un timer enlazado a las señales de reloj del procesador

	competidores := []Competidor{
		{"Competidor 1", rand.Intn(10) + 1},
		{"Competidor 2", rand.Intn(10) + 1}, // Velocidades entre el 1 y 10
		{"Competidor 3", rand.Intn(10) + 1},
		{"Competidor 4", rand.Intn(10) + 1},
	}

	// Creamos un canal para representar la pista
	pista := make(chan int)

	//Nos aseguramos de cerrar el canal cuando se termina la carrera
	defer close(pista)

	//Waitgroup para esperar la finalizacion de todas las go routines
	var wg sync.WaitGroup
	wg.Add(len(competidores))

	//Iniciamos una goroutine para cada competidor

	for _, competidor := range competidores {
		go correrCompetidor(competidor, pista, &wg)
	}

	/*Monitoreamos el progreso de los competidores y definimos
	el ganador*/

/*for {
		select {
		case avance := <-pista:
			fmt.Printf("Avance: %d\n", avance)
		case <-time.After(1 * time.Second):
			fmt.Println("No hay avance en 1 segundo...")

		}

		//verificamos si alguien gana la carrera

		select {
		case <-pista:
			fmt.Println("¡Tenemos un ganador!")
			wg.Wait()
			return
		default:
			//No hay ganador y se continua la carrera
		}
	}
}*/

/*

goroutine 19 [chan send]:

Esto quiere decir de que en la unidad de memoria 19 se envio informacion via un canal.

Para que no se vayan a asustar si les sale algo parecido a la hora de desarrollar.

Si, tengamos en cuenta de que cuando trabajamos con goroutines y canales, estamos trabajando con unidades aisladas de memoria volatil del procesador.


Para todos quedo entendido ?

Quieren que resolvamos otro ?


ok

cual quieren ?


Ok ya vamos con ello

*/

//Primero desarrollemos los imports
