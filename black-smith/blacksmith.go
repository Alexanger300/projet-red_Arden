package blacksmith

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Arden/character"
	"github.com/Alexanger300/projet-red_Arden/inventory"
)

// Menu principal du forgeron
func Welcome(player *character.Character, inv *inventory.Inventory) {
	fmt.Println("=== Bienvenue chez le forgeron ===")
	fmt.Println("Je peux forger des armes pour vous ou améliorer votre équipement.")

	for {
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Forger une nouvelle arme")
		fmt.Println("2. Améliorer votre équipement")
		fmt.Println("3. Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			forgeWeapon(player, inv)
		case 2:
			improveEquipment(player, inv)
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, essayez encore.")
		}
	}
}

// Forge d'armes
func forgeWeapon(player *character.Character, inv *inventory.Inventory) {
	fmt.Println("\n--- Forge d'Armes ---")
	fmt.Println("1. Épée sacrée (Paladin) [2 peaux de gobelins, 1 lingot de fer + 10 or]")
	fmt.Println("2. Gantelets colossaux (Géant) [2 cuirs de sangliers, 1 fourrure de loup + 10 or]")
	fmt.Println("3. Grimoire ancien (Mage) [1 parchemin ancien, 2 cristaux magiques + 10 or]")
	fmt.Println("4. Bâton de vie (Guérisseur) [3 branches d'arbre, 1 pierre de vie + 10 or]")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if player.Wallet.Spend(10) && inv.HasItem("Peau de gobelin", 2) && inv.HasItem("Lingot de fer", 1) {
			inv.RemoveItem("Peau de gobelin", 2)
			inv.RemoveItem("Lingot de fer", 1)
			player.Weapon = "Épée sacrée"
			fmt.Println("Vous avez forgé une Épée sacrée !")
		} else {
			fmt.Println("Pas assez d’or ou de ressources.")
		}
	case 2:
		if player.Wallet.Spend(10) && inv.HasItem("Cuir de sanglier", 2) && inv.HasItem("Fourrure de loup", 1) {
			inv.RemoveItem("Cuir de sanglier", 2)
			inv.RemoveItem("Fourrure de loup", 1)
			player.Weapon = "Gantelets colossaux"
			fmt.Println("Vous avez forgé des Gantelets colossaux !")
		} else {
			fmt.Println("Pas assez d’or ou de ressources.")
		}
	case 3:
		if player.Wallet.Spend(10) && inv.HasItem("Parchemin ancien", 1) && inv.HasItem("Cristal magique", 2) {
			inv.RemoveItem("Parchemin ancien", 1)
			inv.RemoveItem("Cristal magique", 2)
			player.Weapon = "Grimoire ancien"
			fmt.Println("Vous avez forgé un Grimoire ancien !")
		} else {
			fmt.Println("Pas assez d’or ou de ressources.")
		}
	case 4:
		if player.Wallet.Spend(10) && inv.HasItem("Branche d'arbre", 3) && inv.HasItem("Pierre de vie", 1) {
			inv.RemoveItem("Branche d'arbre", 3)
			inv.RemoveItem("Pierre de vie", 1)
			player.Weapon = "Bâton de vie"
			fmt.Println("Vous avez forgé un Bâton de vie !")
		} else {
			fmt.Println("Pas assez d’or ou de ressources.")
		}
	case 0:
		return
	default:
		fmt.Println("Choix invalide, essaie encore.")
	}
}

// Amélioration équipement
func improveEquipment(player *character.Character, inv *inventory.Inventory) {
	fmt.Println("\n--- Amélioration d'Équipement ---")
	fmt.Println("1. Améliorer une armure (5 or + ressource simple)")
	fmt.Println("2. Améliorer une arme (10 or + ressource rare)")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if player.Wallet.Spend(5) && inv.HasItem("Ressource simple", 1) {
			inv.RemoveItem("Ressource simple", 1)
			player.Def += 5
			fmt.Println("Votre armure a été améliorée ! DEF +5")
		} else {
			fmt.Println("Pas assez d’or ou de ressources.")
		}
	case 2:
		if player.Wallet.Spend(10) && inv.HasItem("Ressource rare", 1) {
			inv.RemoveItem("Ressource rare", 1)
			player.Atk += 5
			fmt.Println("Votre arme a été améliorée ! ATK +5")
		} else {
			fmt.Println("Pas assez d’or ou de ressources.")
		}
	case 0:
		return
	default:
		fmt.Println("Choix invalide, essaie encore.")
	}
}
