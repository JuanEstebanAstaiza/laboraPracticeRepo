package main

/*import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Función para producir/fabricar numeros aleatorios y enviarlos en un canal

func productor(numeros chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		numero := rand.Intn(100)
		numeros <- numero
		fmt.Printf("Productor: Enviado número %d\n", numero)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) //Simulamos una produccion irregular
	}

	close(numeros)
}

//Funcion  consumidora que recibe numeros del canal y los procesa

func consumidor(numeros <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for numero := range numeros {

		fmt.Printf("Consumidor: Recibido número %d\n", numero)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//Creamos el canal de comunicacion entre productores y consumidores

	numeros := make(chan int)

	// Waitgropu para esperar la finalizacion de las goroutines
	var wg sync.WaitGroup

	// Iniciamos las goroutines de productor y consumidor

	wg.Add(2)
	go productor(numeros, &wg)
	go consumidor(numeros, &wg)

	//Esperamos a que terminen

	wg.Wait()
}*/

/*

Dime.

No importa si retornan, recordemos que el concepto go routine es para funciones que tengan que trabajar de manera modular, si observas ellas retornan los numeros que se reciben

Si quieren retornar y almacenar se puede almacenar via estructura de datos, se puede guardar en arrays por lo general ya que los datos de un canal son de un solo tipo.


sip

De hecho cuando ustedes hacen consultas a una API, todos los datos viajan por un canal predeterminado, ustedes los almacenan de diferentes formas siempre y cuando se mantenga la integridad de los mismos
Si, se reciben los datos que se enviaron por el canal llamado numeros

No, solo con que este declarado el canal ya se envian

Ahora si les quedo claro ?

Perfecto muchachos, es hora de ponerse manos a la obra, demos la clase cerrada en este momento, ya lo subo al repositorio de una vez, y, no se preocupen si les falta terminar el trabajo de hoy les doy un plazo maximo hasta el domingo.

No te preocupes

Bueno muchachos

Pasen linda noche, trabajen bastante, me alegra que hayan entendido los conceptos de esta prueba tecnica

Igual ya les queda subida

Cualquier duda me escriben al Slack o Whatsapp


*/
