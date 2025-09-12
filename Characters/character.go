package main

import (
	"fmt"
)

type Character struct {
	name  string
	class string
}

func initCharacter() {
	confirmed := false
	var main_character Character
	var choice_number int
	var YesNo string
	//Nom Du Personnage//
	fmt.Print("Entrez le nom du personnage :")
	_, err := fmt.Scan(&main_character.name)
	if err != nil {
		fmt.Println("Error", err) //Renvoie Une erreur
	}

	fmt.Printf("Voici le nom de votre personnage %s", main_character.name)
	fmt.Println()
	for !confirmed {
		//Classe du Personnage//
		fmt.Println("Quelle Classe voulez-vous ? ")
		fmt.Println("1: Paladin ")
		fmt.Println("2:Géant  ")
		fmt.Println("3:Mage ")
		fmt.Println("4:Guérisseur ")
		fmt.Scan(&choice_number)
		switch choice_number {
		case 1:
			fmt.Println("Nourri par la foi et la justice, le Paladin porte une épée sacrée. Il protège les faibles et lutte contre les hérétiques. Sa force n’est pas seulement dans ses bras, mais dans sa croyance inébranlable.")
			fmt.Println("Confirmez-vous votre choix ? (Oui/Non)")
			fmt.Scan(&YesNo)
			if YesNo == "Oui" {
				fmt.Println("Vous êtes Paladin")
				main_character.class = "Paladin"
				confirmed = true
			}
		case 2:
			fmt.Println("Né des montagnes, le Géant n’a jamais connu la peur.Ses poings sont des armes, son corps un mur. Il n’obéit à personne, mais quand il choisit un camp, il le défend jusqu’à la mort.")
			fmt.Println("Confirmez-vous votre choix ? (Oui/Non)")
			fmt.Scan(&YesNo)
			if YesNo == "Oui" {
				fmt.Println("Vous êtes Géant")
				main_character.class = "Géant"
				confirmed = true
			}
		case 3:
			fmt.Println("Le Mage a délaissé les lames et les boucliers. Dans son vieux grimoire sommeillent des flammes, des éclairs et des ombres interdites.Le monde le craint, car là où il passe, la réalité elle-même se plie.")
			fmt.Println("Confirmez-vous votre choix ? (Oui/Non)")
			fmt.Scan(&YesNo)
			if YesNo == "Oui" {
				fmt.Println("Vous êtes Mage")
				main_character.class = "Mage"
				confirmed = true
			}
		case 4:
			fmt.Println("Le Guérisseur n’est pas un guerrier, mais un gardien de vie. Son bâton n’apporte pas la mort mais la lumière.Beaucoup se moquent de lui, jusqu’au jour où ses soins sauvent tout un bataillon.")
			fmt.Println("Confirmez-vous votre choix ? (Oui/Non)")
			fmt.Scan(&YesNo)
			if YesNo == "Oui" {
				fmt.Println("Vous êtes Guérisseur")
				main_character.class = "Guérisseur"
				confirmed = true
			}
		default:
			fmt.Println("PasDeClasse")
		}

		fmt.Printf("-->Votre nom est %s", main_character.name)
		fmt.Println(" ")
		fmt.Printf("-->Votre classe est: %s", main_character.class)
	}
}
func main() {
	initCharacter()
}
