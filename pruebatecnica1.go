package main

import "fmt"

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

func main() {
	//se define el n

	n := 100

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
}
