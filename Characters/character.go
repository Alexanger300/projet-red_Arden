package main

import (
	"fmt"
)

type Character struct {
	name  string
	class string
}

func initCharacter() {
	var main_character Character
	var choice_number int
	//Nom Du Personnage//
	fmt.Print("Entrez le nom du personnage :")
	_, err := fmt.Scan(&main_character.name)
	if err != nil {
		fmt.Println("Error", err) //Renvoie Une erreur
	}

	fmt.Printf("Voici le nom de votre personnage %s", main_character.name)
	fmt.Println()

	//Classe du Personnage//
	fmt.Println("Quelle Classe voulez-vous ? ")
	fmt.Println("1: Paladin ")
	fmt.Println("2:Géant  ")
	fmt.Println("3:Mage ")
	fmt.Println("4:Guérisseur ")
	fmt.Scan(&choice_number)
	switch choice_number {
	case 1:
		fmt.Println("Vous êtes Paladin")
		main_character.class = "Paladin"
	case 2:
		fmt.Println("Vous êtes Géant")
		main_character.class = "Géant"
	case 3:
		fmt.Println("Vous êtes Mage")
		main_character.class = "Mage"
	case 4:
		fmt.Println("Vous êtes Guérisseur")
		main_character.class = "Guérisseur"
	default:
		fmt.Println("PasDeClasse")
	}
	fmt.Printf("-->Votre nom est %s", main_character.name)
	fmt.Println(" ")
	fmt.Printf("-->Votre classe est: %s", main_character.class)

}

func main() {
	initCharacter()
}
