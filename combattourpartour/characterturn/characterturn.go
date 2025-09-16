package main

import (
	"fmt"
	"time"
) //A importer -->Character //Menu des sorts

func characterTurn(character *Character) {
	fmt.Println("C'est au tour de", character.Name)
	time.Sleep(1 * time.Second)
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1. Attaque Basique :")
	fmt.Println("2. Attaque Spéciale :")
	fmt.Println("Inventaire :")
	var choice int
	fmt.Scan(&choice)
	for {
		switch choice {
		case 1:
			fmt.Println(character.Name, "utilise son attaque basique :")

		case 2:
			//Ouvre un menu de sorts
		case 3:
			//Ouvre un menu inventaire
		default:
			fmt.Println("Choix invalide, vous ratez votre tour !")

		}
	}

}
