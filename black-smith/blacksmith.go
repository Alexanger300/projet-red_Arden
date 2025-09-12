package blacksmith

import "fmt"

func Welcome() {
	// Bienvenue chez le forgeron
	fmt.Println("Bienvenue chez le forgeron !")
	fmt.Println("Je peux forger des armes pour vous ou améliorer votre équipement.")

	for {
		// Menu principal du forgeron
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Forger une nouvelle arme")
		fmt.Println("2. Améliorer votre équipement")
		fmt.Println("3. Quitter")

		var choice int
		// Lire le choix de l'utilisateur
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		// Ajouter les cas pour chaque option
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
	// Menu de forge d'armes
	fmt.Println("\n--- Forge d'Armes ---")
	fmt.Println("Choisissez le type d'arme à forger :")
	fmt.Println("1. Épée sacrée (Paladin) [2 peaux de gobelins, 1 lingot de fer + 10 pièces d'or]")
	fmt.Println("2. Gantelets colossaux (Géant) [2 cuirs de sangliers, 1 fourrure de loup + 10 pièces d'or]")
	fmt.Println("3. Grimoire ancien (Mage) [1 parchemin ancien, 2 cristaux magiques + 10 pièces d'or]")
	fmt.Println("4. Bâton de vie (Guérisseur) [3 branches d'arbre, 1 pierre de vie + 10 pièces d'or]")
	fmt.Println("0. Retour")

	var choice int
	// Lire le choix de l'utilisateur
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	// Ajouter les cas pour chaque arme
	case 1:
		fmt.Println("Vous avez choisi de forger une Épée sacrée !")
	case 2:
		fmt.Println("Vous avez choisi de forger des Gantelets colossaux !")
	case 3:
		fmt.Println("Vous avez choisi de forger un Grimoire ancien !")
	case 4:
		fmt.Println("Vous avez choisi de forger un Bâton de vie !")
	case 0:
		Welcome()
	default:
		fmt.Println("Choix invalide, essaie encore.")
		forgeWeapon()
	}
}

func improveEquipment() {
	// Menu d'amélioration d'équipement
	fmt.Println("\n--- Amélioration d'Équipement ---")
	fmt.Println("1. Améliorer une armure (5 or + ressource simple)")
	fmt.Println("2. Améliorer une arme (10 or + ressource rare)")
	fmt.Println("0. Retour")

	var choice int
	// Lire le choix de l'utilisateur
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	// Ajouter les cas pour chaque amélioration
	case 1:
		fmt.Println("Vous avez choisi d'améliorer une armure !")
	case 2:
		fmt.Println("Vous avez choisi d'améliorer une arme !")
	case 0:
		Welcome()
	default:
		fmt.Println("Choix invalide, essaie encore.")
		improveEquipment()
	}
}
