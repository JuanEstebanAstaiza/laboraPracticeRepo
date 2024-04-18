package main

//football players technical test

/*
import (
	"fmt"
	"strings"
)

type player struct {
	Name       string
	Identifier int
}

func main() {
	//Lista jugadores

	players := []string{"John Smith", "David Jhonson", "Michael Garcia", "Sara Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sofia Lee", "Lucas Oliviera", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}

	newPlayers := []string{"New player 1", "New player 2", "New player", "New player 3", "New player 4", "New player 5", "New player 6", "New player 7"}

	players = players[7:]

	for x := 0; x < len(players); x++ {
		fmt.Println(players[x])
	}

	//agregar
	for _, name := range newPlayers {
		players = append(players, name)
	}

	// Mostrar lista completa

	fmt.Println("lista completa de jugadores: ")
	for i, player := range players {
		fmt.Println("\n", i+1001, sanitizeName(player))
	}

}

func sanitizeName(name string) string {

	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}

		return -1
	}, name)

}*/

//fizz buzz test
/*
func main() {
	//se define el n

	n := 10000

	//imprimimos el algoritmo

	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

		if i < n {
			fmt.Println(", ")
		}
	}
}*/

/*var (
	adminUsername      = "admin"
	adminPassword      = "root"
	supervisorUsername = "seeker223"
	supervisorPassword = "seekr"
	adminActions       = []string{"Crear laborer", "Eliminar laborer", "Generar archivos"}
	laborers           []string
)

func main() {
	for {
		fmt.Println("Bienvenido a Labora")
		fmt.Println("1. Iniciar sesión como administrador")
		fmt.Println("2. Iniciar sesión como supervisor")
		fmt.Println("3. Salir")

		var choice int
		fmt.Println("Ingrese su opción: ")
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
			fmt.Println("Opción invalida. Porfavor, intente de nuevo.")
		}
	}
}

func adminLogin() {
	fmt.Println("Ingrese las creadeciales de Administrador")
	fmt.Println("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == adminUsername && password == adminPassword {
		fmt.Println("Inicio de sesion exitoso como Administrador! ")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")

	}
}

func supervisorLogin() {
	fmt.Println("Ingrese las credenciales de Supervisor")
	fmt.Println("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Contraseña: ")
	var password string
	fmt.Scanln(password)

	if username == supervisorUsername && password == supervisorPassword {
		fmt.Println("Inicio de sesión exitoso como Supervisor")
		//supervisorMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intentelo de nuevo.")

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
		fmt.Println("Seleccione su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Acción: Crear laborer")
			createLaborer()
		case 2:
			fmt.Println("Acción: Eliminar laborer")
			deleteLaborer()
		case 3:
			fmt.Println("Acción: Generar archivo de texto personalizado")
			generateTextFile()
		case 0:
			fmt.Println("Cerrando sesión de Administrador")
			return
		default:
			fmt.Println("Opción invalida. Por favor, intente de nuevo")
		}
	}
}

func createLaborer() {
	laborers = append(laborers, fmt.Sprintf("laborer %d", len(laborers)+1))
	fmt.Println("Laborer creado exitosamente")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}
}

func deleteLaborer() {
	if len(laborers) == 0 {
		fmt.Println("No hay laborers para eliminar")
		return
	}

	laborers = laborers[:len(laborers)-1]
	fmt.Println("laborer eliminado exitosamente")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}
}

func generateTextFile() {
	fileName := fmt.Sprintf("mensajes_%s.txt", time.Now().Format("20060102_152032"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear archivo: %v\n", err)
		return
	}
	defer file.Close()

	for _, laborer := range laborers {
		_, err := file.WriteString(fmt.Sprintf("%s\n", laborer))
		if err != nil {
			fmt.Printf("Error al escribir el archivo %v\n", err)
			return
		}
	}

	fmt.Printf("Archivo %s creado exitosamente. \n", fileName)
}*/
