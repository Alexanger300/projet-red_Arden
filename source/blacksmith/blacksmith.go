package blacksmith

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/inventory"
)

// Fonction utilitaire pour afficher les ressources demandées avec ce que possède le joueur
func displayRequirement(player *character.Character, item string, needed int) string {
	have := inventory.CountItem(item)
	return fmt.Sprintf("%s (%d/%d)", item, have, needed)
}

func Welcome(player *character.Character) {
	fmt.Println("=== Bienvenue chez le forgeron ===")
	fmt.Println("Je peux forger de l’équipement pour vous ou améliorer ce que vous possédez.")

	for {
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Forger un équipement")
		fmt.Println("2. Améliorer votre équipement")
		fmt.Println("3. Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			forgeEquipment(player)
		case 2:
			improveEquipment(player)
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("❌ Choix invalide, essayez encore.")
		}
	}
}

// Forge
func forgeEquipment(player *character.Character) {
	fmt.Println("\n--- Forge ---")
	fmt.Println("Voici les recettes disponibles :")

	var choice int

	switch player.Class {
	//Paladin
	case "Paladin":
		fmt.Println("1. Épée sacrée ⚔️ →", displayRequirement(player, "Peau de gobelin", 2)+",",
			displayRequirement(player, "Lingot de fer", 1)+",", fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Casque de paladin 🪖 →", displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Branche d'arbre", 1)+",", fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Armure bénie 🛡️ →", displayRequirement(player, "Lingot de fer", 3)+",",
			displayRequirement(player, "Cristal de vie", 1)+",", fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Jambières lourdes 🦵 →", displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Pierre de vie", 1)+",", fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Spend(10) && inventory.HasItem("Peau de gobelin", 2) && inventory.HasItem("Lingot de fer", 1) {
				inventory.RemoveItem("Peau de gobelin", 2)
				inventory.RemoveItem("Lingot de fer", 1)
				player.Weapon = "Épée sacrée"
				fmt.Println("✅ Vous avez forgé une Épée sacrée !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Spend(8) && inventory.HasItem("Lingot de fer", 2) && inventory.HasItem("Branche d'arbre", 1) {
				inventory.RemoveItem("Lingot de fer", 2)
				inventory.RemoveItem("Branche d'arbre", 1)
				player.Equip.Head.Name = "Casque de paladin"
				fmt.Println("✅ Vous avez forgé un Casque de paladin !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Spend(15) && inventory.HasItem("Lingot de fer", 3) && inventory.HasItem("Cristal de vie", 1) {
				inventory.RemoveItem("Lingot de fer", 3)
				inventory.RemoveItem("Cristal de vie", 1)
				player.Equip.Body.Name = "Armure bénie"
				fmt.Println("✅ Vous avez forgé une Armure bénie !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Spend(12) && inventory.HasItem("Lingot de fer", 2) && inventory.HasItem("Pierre de vie", 1) {
				inventory.RemoveItem("Lingot de fer", 2)
				inventory.RemoveItem("Pierre de vie", 1)
				player.Equip.Legs.Name = "Jambières lourdes"
				fmt.Println("✅ Vous avez forgé des Jambières lourdes !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}

	// Géant
	case "Géant":
		fmt.Println("1. Gantelets colossaux 🪓 →", displayRequirement(player, "Cuir de sanglier", 2)+",",
			displayRequirement(player, "Fourrure de loup", 1)+",", fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Heaume massif 🪖 →", displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Cuir de sanglier", 1)+",", fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Plastron de colosse 🛡️ →", displayRequirement(player, "Lingot de fer", 3)+",",
			displayRequirement(player, "Pierre de vie", 1)+",", fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Jambières de pierre 🦵 →", displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Fourrure de loup", 1)+",", fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Spend(10) && inventory.HasItem("Cuir de sanglier", 2) && inventory.HasItem("Fourrure de loup", 1) {
				inventory.RemoveItem("Cuir de sanglier", 2)
				inventory.RemoveItem("Fourrure de loup", 1)
				player.Weapon = "Gantelets colossaux"
				fmt.Println("✅ Vous avez forgé des Gantelets colossaux !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Spend(8) && inventory.HasItem("Lingot de fer", 2) && inventory.HasItem("Cuir de sanglier", 1) {
				inventory.RemoveItem("Lingot de fer", 2)
				inventory.RemoveItem("Cuir de sanglier", 1)
				player.Equip.Head.Name = "Heaume massif"
				fmt.Println("✅ Vous avez forgé un Heaume massif !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Spend(15) && inventory.HasItem("Lingot de fer", 3) && inventory.HasItem("Pierre de vie", 1) {
				inventory.RemoveItem("Lingot de fer", 3)
				inventory.RemoveItem("Pierre de vie", 1)
				player.Equip.Body.Name = "Plastron de colosse"
				fmt.Println("✅ Vous avez forgé un Plastron de colosse !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Spend(12) && inventory.HasItem("Lingot de fer", 2) && inventory.HasItem("Fourrure de loup", 1) {
				inventory.RemoveItem("Lingot de fer", 2)
				inventory.RemoveItem("Fourrure de loup", 1)
				player.Equip.Legs.Name = "Jambières de pierre"
				fmt.Println("✅ Vous avez forgé des Jambières de pierre !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}

	// Mage
	case "Mage":
		fmt.Println("1. Grimoire ancien 📖 →", displayRequirement(player, "Parchemin ancien", 1)+",",
			displayRequirement(player, "Cristal magique", 2)+",", fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Chapeau mystique 🎩 →", displayRequirement(player, "Cristal magique", 1)+",",
			displayRequirement(player, "Branche d'arbre", 2)+",", fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Robe enchantée 🧥 →", displayRequirement(player, "Cristal magique", 2)+",",
			displayRequirement(player, "Pierre de vie", 1)+",", fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Bottes de lévitation 👢 →", displayRequirement(player, "Cristal magique", 1)+",",
			displayRequirement(player, "Peau de gobelin", 2)+",", fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Spend(10) && inventory.HasItem("Parchemin ancien", 1) && inventory.HasItem("Cristal magique", 2) {
				inventory.RemoveItem("Parchemin ancien", 1)
				inventory.RemoveItem("Cristal magique", 2)
				player.Weapon = "Grimoire ancien"
				fmt.Println("✅ Vous avez forgé un Grimoire ancien !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Spend(8) && inventory.HasItem("Cristal magique", 1) && inventory.HasItem("Branche d'arbre", 2) {
				inventory.RemoveItem("Cristal magique", 1)
				inventory.RemoveItem("Branche d'arbre", 2)
				player.Equip.Head.Name = "Chapeau mystique"
				fmt.Println("✅ Vous avez forgé un Chapeau mystique !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Spend(15) && inventory.HasItem("Cristal magique", 2) && inventory.HasItem("Pierre de vie", 1) {
				inventory.RemoveItem("Cristal magique", 2)
				inventory.RemoveItem("Pierre de vie", 1)
				player.Equip.Body.Name = "Robe enchantée"
				fmt.Println("✅ Vous avez forgé une Robe enchantée !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Spend(12) && inventory.HasItem("Cristal magique", 1) && inventory.HasItem("Peau de gobelin", 2) {
				inventory.RemoveItem("Cristal magique", 1)
				inventory.RemoveItem("Peau de gobelin", 2)
				player.Equip.Legs.Name = "Bottes de lévitation"
				fmt.Println("✅ Vous avez forgé des Bottes de lévitation !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}

	// Guérisseur
	case "Guérisseur":
		fmt.Println("1. Bâton de vie 🌿 →", displayRequirement(player, "Branche d'arbre", 3)+",",
			displayRequirement(player, "Pierre de vie", 1)+",", fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Capuche de prêtre 🧢 →", displayRequirement(player, "Branche d'arbre", 2)+",",
			displayRequirement(player, "Cristal de vie", 1)+",", fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Robe de lumière ✨ →", displayRequirement(player, "Cristal de vie", 2)+",",
			displayRequirement(player, "Pierre de vie", 1)+",", fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Sandales bénies 👡 →", displayRequirement(player, "Branche d'arbre", 2)+",",
			displayRequirement(player, "Peau de gobelin", 1)+",", fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")

		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Spend(10) && inventory.HasItem("Branche d'arbre", 3) && inventory.HasItem("Pierre de vie", 1) {
				inventory.RemoveItem("Branche d'arbre", 3)
				inventory.RemoveItem("Pierre de vie", 1)
				player.Weapon = "Bâton de vie"
				fmt.Println("✅ Vous avez forgé un Bâton de vie !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Spend(8) && inventory.HasItem("Branche d'arbre", 2) && inventory.HasItem("Cristal de vie", 1) {
				inventory.RemoveItem("Branche d'arbre", 2)
				inventory.RemoveItem("Cristal de vie", 1)
				player.Equip.Head.Name = "Capuche de prêtre"
				fmt.Println("✅ Vous avez forgé une Capuche de prêtre !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Spend(15) && inventory.HasItem("Cristal de vie", 2) && inventory.HasItem("Pierre de vie", 1) {
				inventory.RemoveItem("Cristal de vie", 2)
				inventory.RemoveItem("Pierre de vie", 1)
				player.Equip.Body.Name = "Robe de lumière"
				fmt.Println("✅ Vous avez forgé une Robe de lumière !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Spend(12) && inventory.HasItem("Branche d'arbre", 2) && inventory.HasItem("Peau de gobelin", 1) {
				inventory.RemoveItem("Branche d'arbre", 2)
				inventory.RemoveItem("Peau de gobelin", 1)
				player.Equip.Legs.Name = "Sandales bénies"
				fmt.Println("✅ Vous avez forgé des Sandales bénies !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}
	}
}

// Amélioration d'équipement
func improveEquipment(player *character.Character) {
	fmt.Println("\n--- Amélioration d'Équipement ---")
	fmt.Println("1. Améliorer une armure (15 or)")
	fmt.Println("2. Améliorer une arme (25 or)")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Vérifier si une armure est équipée
		if player.Equip.Head.Name == "" && player.Equip.Body.Name == "" && player.Equip.Legs.Name == "" {
			fmt.Println("❌ Vous n'avez aucune armure équipée à améliorer.")
			return
		}
		if player.Wallet.Spend(15) {
			player.Def += 5
			fmt.Println("✅ Votre armure a été améliorée ! DEF +5")
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}

	case 2:
		// Vérifier si une arme est équipée
		if player.Weapon == "" {
			fmt.Println("❌ Vous n'avez aucune arme forgée à améliorer.")
			return
		}
		if player.Wallet.Spend(25) {
			player.Atk += 5
			fmt.Printf("✅ Votre arme %s a été améliorée ! ATK +5\n", player.Weapon)
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}

	case 0:
		return
	default:
		fmt.Println("❌ Choix invalide, essaie encore.")
	}
}
