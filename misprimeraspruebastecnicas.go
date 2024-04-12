package main

//prueba tecnica de jugadores de football
/*
import (
	"fmt"
	"strings"
)

type Player struct {
	Name       string
	Identifier int
}

func main() {
	// Lista de jugadores actuales
	players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}

	// Registro de nuevos jugadores
	newPlayers := []string{"New Player 1", "New Player 2", "New Player 3", "New Player 4", "New Player 5", "New Player 6", "New Player 7"}

	// Eliminar los primeros 7 jugadores de la liga anterior
	players = players[7:]

	// Agregar nuevos jugadores
	for _, name := range newPlayers {
		players = append(players, name)
	}

	// Mostrar lista completa de jugadores
	fmt.Println("Lista completa de jugadores:")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1001, sanitizeName(player))
	}
}

// Función para eliminar caracteres especiales y números del nombre del jugador
func sanitizeName(name string) string {
	// Eliminar caracteres especiales y números
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1
	}, name)
}
*/
//prueba tecnica fizz buzz
/*
import "fmt"

func main() {
	// Definir el valor de n
	n := 50

	// Imprimir los números del 1 al n con las reglas de FizzBuzz
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}

		// Agregar coma y espacio para separar los números
		if i < n {
			fmt.Print(", ")
		}
	}
}
*/
/*
//prueba tecnica packages

import (
	"fmt"
	"os"
	"time"
)

var (
	adminUsername      = "admin"
	adminPassword      = "root"
	supervisorUsername = "seeker"
	supervisorPassword = "seekr"

	adminActions  = []string{"Crear laborer", "Eliminar laborer", "Generar archivo de texto con mensajes personalizados"}
	laborers      []string
	actionCounter = make(map[string]int)
)

func main() {
	for {
		fmt.Println("Bienvenido a Labora")
		fmt.Println("1. Iniciar sesión como Administrador")
		fmt.Println("2. Iniciar sesión como Supervisor")
		fmt.Println("3. Salir")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			adminLogin()
		case 2:
			supervisorLogin()
		case 3:
			fmt.Println("¡Hasta luego!")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func adminLogin() {
	fmt.Println("Ingrese las credenciales de Administrador")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == adminUsername && password == adminPassword {
		fmt.Println("Inicio de sesión exitoso como Administrador")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")
	}
}

func supervisorLogin() {
	fmt.Println("Ingrese las credenciales de Supervisor")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == supervisorUsername && password == supervisorPassword {
		fmt.Println("Inicio de sesión exitoso como Supervisor")
		supervisorMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")
	}
}

func adminMenu() {
	for {
		fmt.Println("Menú de Administrador")
		for i, action := range adminActions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createLaborer()
		case 2:
			deleteLaborer()
		case 3:
			fmt.Println("Acción: Generar archivo de texto con mensajes personalizados")
			generateTextFile()
		case 0:
			fmt.Println("Cerrando sesión de Administrador")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func createLaborer() {
	laborers = append(laborers, fmt.Sprintf("laborer %d", len(laborers)+1))
	fmt.Println("Laborer creado exitosamente.")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}

}

func deleteLaborer() {
	if len(laborers) == 0 {
		fmt.Println("No hay laborers para eliminar.")
		return
	}

	laborers = laborers[:len(laborers)-1]
	fmt.Println("Laborer eliminado exitosamente.")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}
}

func supervisorMenu() {
	for {
		fmt.Println("Menú de Supervisor")
		fmt.Println("1. Crear registro de administrador")
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Acción: Crear registro de administrador")
			createAdminRecord()
		case 0:
			fmt.Println("Cerrando sesión de Supervisor")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func generateTextFile() {
	fileName := fmt.Sprintf("mensajes_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir los mensajes en el archivo
	for _, laborer := range laborers {
		_, err := file.WriteString(fmt.Sprintf("%s\n", laborer))
		if err != nil {
			fmt.Printf("Error al escribir en el archivo: %v\n", err)
			return
		}
	}

	fmt.Printf("Archivo %s creado exitosamente.\n", fileName)
}

func createAdminRecord() {
	fileName := fmt.Sprintf("registro_admin_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo de registro: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir los mensajes en el archivo
	for _, action := range adminActions {
		_, err := file.WriteString(fmt.Sprintf("%s: %d\n", action, actionCounter[action]))
		if err != nil {
			fmt.Printf("Error al escribir en el archivo de registro: %v\n", err)
			return
		}
	}

	fmt.Printf("Archivo de registro %s creado exitosamente.\n", fileName)
}


*/
