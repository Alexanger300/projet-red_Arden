package main

import "fmt"

func Welcome() {
	// Bienvenue chez le forgeron
	fmt.Println("Bienvenue chez le forgeron !")
	fmt.Println("Je peux forger des armes pour vous ou améliorer votre équipement.")

	for {
		// Menu des choix
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Forger une nouvelle arme")
		fmt.Println("2. Améliorer votre équipement")
		fmt.Println("3. Quitter")

		var choice int
		// Lecture du choix de l'utilisateur
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			forgeWeapon()
			return
		case 2:
			improveEquipment()
			return
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, essayez encore.")
		}
	}
}

func forgeWeapon() {
	fmt.Println("Vous avez choisi de forger une nouvelle arme.")
	// Ajoutez ici la logique de forge d'arme
}

func improveEquipment() {
	fmt.Println("Vous avez choisi d'améliorer votre équipement.")
	// Ajoutez ici la logique d'amélioration d'équipement
}

func main() {
	Welcome()
}
