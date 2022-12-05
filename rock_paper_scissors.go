package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var sceltaUtente int
	for {
		fmt.Print("1.Sasso", "\n", "2.Carta", "\n", "3.Forbice", "\n", "0.Fine")
		fmt.Println()
		fmt.Scan(&sceltaUtente)
		if sceltaUtente == 1 {
			if ComputerChoice() == 3 {
				fmt.Println("Hai vinto!")
			} else if ComputerChoice() == 2 {
				fmt.Println("Pareggio.")
			} else if ComputerChoice() == 1 {
				fmt.Println("Hai perso!")
			}
		} else if sceltaUtente == 2 {
			if ComputerChoice() == 2 {
				fmt.Println("Hai vinto")
			} else if ComputerChoice() == 1 {
				fmt.Println("Pareggio.")
			} else if ComputerChoice() == 3 {
				fmt.Println("Hai perso!")
			}
		} else if sceltaUtente == 3 {
			if ComputerChoice() == 1 {
				fmt.Println("Hai vinto!")
			} else if ComputerChoice() == 3 {
				fmt.Println("Pareggio.")
			} else if ComputerChoice() == 2 {
				fmt.Println("Hai perso!")
			}
		} else if sceltaUtente == 4 {
			fmt.Println("Le vinci tutte!")
		} else if sceltaUtente == 0 {
			break
		}
	}
}

func ComputerChoice() int {
	min := 1
	max := 4
	rand.Seed(time.Now().UnixNano())
	randomNumber := (rand.Intn(max-min) + min)
	switch randomNumber {
	case '1':
		fmt.Println("carta")
	case '2':
		fmt.Println("sasso")
	case '3':
		fmt.Println("forbice")
	}
	return randomNumber
}
