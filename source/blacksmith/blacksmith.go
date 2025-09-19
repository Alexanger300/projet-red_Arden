package blacksmith

import (
	"fmt"
	"time"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/equipment"
)

// Assure que la map des armes est initialisée.

func ensureWeaponsMap(player *character.Character) {
	if player.Weapons == nil {
		player.Weapons = make(map[string]equipment.Equipment)
	}
}

// Retire une quantité d’un item. Supprime la clé si quantité = 0.
func removeItem(player *character.Character, item string, qty int) {
	if qty <= 0 {
		return
	}
	if player.Inventory[item] >= qty {
		player.Inventory[item] -= qty
		if player.Inventory[item] == 0 {
			delete(player.Inventory, item)
		}
	}
}

func hasItem(player *character.Character, item string, qty int) bool {
	return player.Inventory[item] >= qty
}

func displayRequirement(player *character.Character, item string, needed int) string {
	have := player.Inventory[item]
	return fmt.Sprintf("%s (%d/%d)", item, have, needed)
}

func equipWeapon(player *character.Character, w equipment.Equipment) {
	ensureWeaponsMap(player)
	player.Weapons[w.Name] = w
	player.Equip.Weapon = w
	player.RecalculateStatsFromEquipment()
}

// Menu Principal du forgeron

func Welcome(player *character.Character) {
	css.Clear()
	text := "=== " + "⚒️ " + " Bienvenue chez le forgeron ==="
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}
	time.Sleep(500 * time.Millisecond)
	text1 := "\nJe peux forger de l'équipement pour vous ou améliorer ce que vous possédez."
	for _, char := range text1 {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)

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
			css.Clear()
			return
		default:
			fmt.Println("❌ Choix invalide, essayez encore.")
		}
	}
}

// 	Forger l'équipement selon la classe

func forgeEquipment(player *character.Character) {
	fmt.Println("\n--- 🔨 Forge ---")
	fmt.Println("Voici les recettes disponibles :")

	var choice int

	switch player.Class {

	// Pour Paladin
	case "Paladin":
		fmt.Println("1. Épée sacrée ⚔️ →",
			displayRequirement(player, "Peau de gobelin", 2)+",",
			displayRequirement(player, "Lingot de fer", 1)+",",
			fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Casque de paladin 🪖 →",
			displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Branche d'arbre", 1)+",",
			fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Armure bénie 🛡️ →",
			displayRequirement(player, "Lingot de fer", 3)+",",
			displayRequirement(player, "Cristal de vie", 1)+",",
			fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Jambières lourdes 🦵 →",
			displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Pierre de vie", 1)+",",
			fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Amount >= 10 &&
				hasItem(player, "Peau de gobelin", 2) &&
				hasItem(player, "Lingot de fer", 1) {
				removeItem(player, "Peau de gobelin", 2)
				removeItem(player, "Lingot de fer", 1)
				player.Wallet.Spend(10)
				weapon := equipment.Equipment{Name: "Épée sacrée", Atk: 5, Slot: "Weapon", Class: "Paladin"}
				equipWeapon(player, weapon)
				fmt.Println("✅ Vous avez forgé une Épée sacrée (+5 ATK) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Amount >= 8 &&
				hasItem(player, "Lingot de fer", 2) &&
				hasItem(player, "Branche d'arbre", 1) {
				removeItem(player, "Lingot de fer", 2)
				removeItem(player, "Branche d'arbre", 1)
				player.Wallet.Spend(8)
				player.Equip.Head = equipment.Equipment{Name: "Casque de paladin", Def: 5, Slot: "Head", Class: "Paladin"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé un Casque de paladin (+5 DEF) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Amount >= 15 &&
				hasItem(player, "Lingot de fer", 3) &&
				hasItem(player, "Cristal de vie", 1) {
				removeItem(player, "Lingot de fer", 3)
				removeItem(player, "Cristal de vie", 1)
				player.Wallet.Spend(15)
				player.Equip.Body = equipment.Equipment{Name: "Armure bénie", Def: 8, HP: 20, Slot: "Body", Class: "Paladin"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé une Armure bénie (+8 DEF, +20 HP) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Amount >= 12 &&
				hasItem(player, "Lingot de fer", 2) &&
				hasItem(player, "Pierre de vie", 1) {
				removeItem(player, "Lingot de fer", 2)
				removeItem(player, "Pierre de vie", 1)
				player.Wallet.Spend(12)
				player.Equip.Legs = equipment.Equipment{Name: "Jambières lourdes", Def: 6, Spd: -1, Slot: "Legs", Class: "Paladin"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé des Jambières lourdes (+6 DEF, -1 SPD) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}

	// Pour Géant
	case "Géant":
		fmt.Println("1. Gantelets colossaux 🪓 →",
			displayRequirement(player, "Cuir de sanglier", 2)+",",
			displayRequirement(player, "Fourrure de loup", 1)+",",
			fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Heaume massif 🪖 →",
			displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Cuir de sanglier", 1)+",",
			fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Plastron de colosse 🛡️ →",
			displayRequirement(player, "Lingot de fer", 3)+",",
			displayRequirement(player, "Pierre de vie", 1)+",",
			fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Jambières de pierre 🦵 →",
			displayRequirement(player, "Lingot de fer", 2)+",",
			displayRequirement(player, "Fourrure de loup", 1)+",",
			fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Amount >= 10 &&
				hasItem(player, "Cuir de sanglier", 2) &&
				hasItem(player, "Fourrure de loup", 1) {
				removeItem(player, "Cuir de sanglier", 2)
				removeItem(player, "Fourrure de loup", 1)
				player.Wallet.Spend(10)
				weapon := equipment.Equipment{Name: "Gantelets colossaux", Atk: 7, Slot: "Weapon", Class: "Géant"}
				equipWeapon(player, weapon)
				fmt.Println("✅ Vous avez forgé des Gantelets colossaux (+7 ATK) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Amount >= 8 &&
				hasItem(player, "Lingot de fer", 2) &&
				hasItem(player, "Cuir de sanglier", 1) {
				removeItem(player, "Lingot de fer", 2)
				removeItem(player, "Cuir de sanglier", 1)
				player.Wallet.Spend(8)
				player.Equip.Head = equipment.Equipment{Name: "Heaume massif", Def: 6, Slot: "Head", Class: "Géant"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé un Heaume massif (+6 DEF) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Amount >= 15 &&
				hasItem(player, "Lingot de fer", 3) &&
				hasItem(player, "Pierre de vie", 1) {
				removeItem(player, "Lingot de fer", 3)
				removeItem(player, "Pierre de vie", 1)
				player.Wallet.Spend(15)
				player.Equip.Body = equipment.Equipment{Name: "Plastron de colosse", Def: 10, HP: 25, Slot: "Body", Class: "Géant"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé un Plastron de colosse (+10 DEF, +25 HP) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Amount >= 12 &&
				hasItem(player, "Lingot de fer", 2) &&
				hasItem(player, "Fourrure de loup", 1) {
				removeItem(player, "Lingot de fer", 2)
				removeItem(player, "Fourrure de loup", 1)
				player.Wallet.Spend(12)
				player.Equip.Legs = equipment.Equipment{Name: "Jambières de pierre", Def: 7, Spd: -2, Slot: "Legs", Class: "Géant"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé des Jambières de pierre (+7 DEF, -2 SPD) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}

	// Pour Mage
	case "Mage":
		fmt.Println("1. Grimoire ancien 📖 →",
			displayRequirement(player, "Parchemin ancien", 1)+",",
			displayRequirement(player, "Cristal magique", 2)+",",
			fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Chapeau mystique 🎩 →",
			displayRequirement(player, "Cristal magique", 1)+",",
			displayRequirement(player, "Branche d'arbre", 2)+",",
			fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Robe enchantée 🧥 →",
			displayRequirement(player, "Cristal magique", 2)+",",
			displayRequirement(player, "Pierre de vie", 1)+",",
			fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Bottes de lévitation 👢 →",
			displayRequirement(player, "Cristal magique", 1)+",",
			displayRequirement(player, "Peau de gobelin", 2)+",",
			fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Amount >= 10 &&
				hasItem(player, "Parchemin ancien", 1) &&
				hasItem(player, "Cristal magique", 2) {
				removeItem(player, "Parchemin ancien", 1)
				removeItem(player, "Cristal magique", 2)
				player.Wallet.Spend(10)
				weapon := equipment.Equipment{Name: "Grimoire ancien", Atk: 4, Mana: 15, Slot: "Weapon", Class: "Mage"}
				equipWeapon(player, weapon)
				fmt.Println("✅ Vous avez forgé un Grimoire ancien (+4 ATK, +15 Mana) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Amount >= 8 &&
				hasItem(player, "Cristal magique", 1) &&
				hasItem(player, "Branche d'arbre", 2) {
				removeItem(player, "Cristal magique", 1)
				removeItem(player, "Branche d'arbre", 2)
				player.Wallet.Spend(8)
				player.Equip.Head = equipment.Equipment{Name: "Chapeau mystique", Def: 3, Mana: 10, Slot: "Head", Class: "Mage"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé un Chapeau mystique (+3 DEF, +10 Mana) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Amount >= 15 &&
				hasItem(player, "Cristal magique", 2) &&
				hasItem(player, "Pierre de vie", 1) {
				removeItem(player, "Cristal magique", 2)
				removeItem(player, "Pierre de vie", 1)
				player.Wallet.Spend(15)
				player.Equip.Body = equipment.Equipment{Name: "Robe enchantée", Def: 4, Mana: 15, Slot: "Body", Class: "Mage"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé une Robe enchantée (+4 DEF, +15 Mana) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Amount >= 12 &&
				hasItem(player, "Cristal magique", 1) &&
				hasItem(player, "Peau de gobelin", 2) {
				removeItem(player, "Cristal magique", 1)
				removeItem(player, "Peau de gobelin", 2)
				player.Wallet.Spend(12)
				player.Equip.Legs = equipment.Equipment{Name: "Bottes de lévitation", Spd: 3, Mana: 5, Slot: "Legs", Class: "Mage"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé des Bottes de lévitation (+3 SPD, +5 Mana) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}

	// Pour Guérisseur
	case "Guérisseur":
		fmt.Println("1. Bâton de vie 🌿 →",
			displayRequirement(player, "Branche d'arbre", 3)+",",
			displayRequirement(player, "Pierre de vie", 1)+",",
			fmt.Sprintf("Or (%d/10)", player.Wallet.Amount))
		fmt.Println("2. Capuche de prêtre 🧢 →",
			displayRequirement(player, "Branche d'arbre", 2)+",",
			displayRequirement(player, "Cristal de vie", 1)+",",
			fmt.Sprintf("Or (%d/8)", player.Wallet.Amount))
		fmt.Println("3. Robe de lumière ✨ →",
			displayRequirement(player, "Cristal de vie", 2)+",",
			displayRequirement(player, "Pierre de vie", 1)+",",
			fmt.Sprintf("Or (%d/15)", player.Wallet.Amount))
		fmt.Println("4. Sandales bénies 👡 →",
			displayRequirement(player, "Branche d'arbre", 2)+",",
			displayRequirement(player, "Peau de gobelin", 1)+",",
			fmt.Sprintf("Or (%d/12)", player.Wallet.Amount))
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if player.Wallet.Spend(10) &&
				hasItem(player, "Branche d'arbre", 3) &&
				hasItem(player, "Pierre de vie", 1) {
				removeItem(player, "Branche d'arbre", 3)
				removeItem(player, "Pierre de vie", 1)
				weapon := equipment.Equipment{Name: "Bâton de vie", Atk: 3, Mana: 10, Slot: "Weapon", Class: "Guérisseur"}
				equipWeapon(player, weapon)
				fmt.Println("✅ Vous avez forgé un Bâton de vie (+3 ATK, +10 Mana) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 2:
			if player.Wallet.Spend(8) &&
				hasItem(player, "Branche d'arbre", 2) &&
				hasItem(player, "Cristal de vie", 1) {
				removeItem(player, "Branche d'arbre", 2)
				removeItem(player, "Cristal de vie", 1)
				player.Equip.Head = equipment.Equipment{Name: "Capuche de prêtre", Def: 2, Mana: 5, Slot: "Head", Class: "Guérisseur"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé une Capuche de prêtre (+2 DEF, +5 Mana) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 3:
			if player.Wallet.Spend(15) &&
				hasItem(player, "Cristal de vie", 2) &&
				hasItem(player, "Pierre de vie", 1) {
				removeItem(player, "Cristal de vie", 2)
				removeItem(player, "Pierre de vie", 1)
				player.Equip.Body = equipment.Equipment{Name: "Robe de lumière", Def: 5, Mana: 10, HP: 10, Slot: "Body", Class: "Guérisseur"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé une Robe de lumière (+5 DEF, +10 Mana, +10 HP) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		case 4:
			if player.Wallet.Spend(12) &&
				hasItem(player, "Branche d'arbre", 2) &&
				hasItem(player, "Peau de gobelin", 1) {
				removeItem(player, "Branche d'arbre", 2)
				removeItem(player, "Peau de gobelin", 1)
				player.Equip.Legs = equipment.Equipment{Name: "Sandales bénies", Spd: 2, Slot: "Legs", Class: "Guérisseur"}
				player.RecalculateStatsFromEquipment()
				fmt.Println("✅ Vous avez forgé des Sandales bénies (+2 SPD) !")
			} else {
				fmt.Println("❌ Pas assez d'or ou de ressources.")
			}
		}
	default:
		fmt.Println("❌ Classe non gérée pour la forge.")
	}
}

// Améliorer l'équipement existant

func improveEquipment(player *character.Character) {
	fmt.Println("\n--- 🔧 Amélioration d'Équipement ---")
	fmt.Println("1. Améliorer une armure (15 or)")
	fmt.Println("2. Améliorer une arme (25 or)")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if player.Equip.Head.Name == "" && player.Equip.Body.Name == "" && player.Equip.Legs.Name == "" {
			fmt.Println("❌ Vous n'avez aucune armure équipée à améliorer.")
			return
		}

		fmt.Println("\nQuelle pièce voulez-vous améliorer ?")
		i := 1
		choices := make(map[int]*equipment.Equipment)

		if player.Equip.Head.Name != "" {
			fmt.Printf("%d. Casque : %s\n", i, player.Equip.Head.Name)
			choices[i] = &player.Equip.Head
			i++
		}
		if player.Equip.Body.Name != "" {
			fmt.Printf("%d. Armure : %s\n", i, player.Equip.Body.Name)
			choices[i] = &player.Equip.Body
			i++
		}
		if player.Equip.Legs.Name != "" {
			fmt.Printf("%d. Jambières : %s\n", i, player.Equip.Legs.Name)
			choices[i] = &player.Equip.Legs
			i++
		}

		var partChoice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&partChoice)

		chosen, ok := choices[partChoice]
		if !ok {
			fmt.Println("❌ Choix invalide.")
			return
		}

		if player.Wallet.Spend(15) {
			chosen.Def += 5
			player.RecalculateStatsFromEquipment()
			fmt.Printf("✅ %s amélioré(e) ! DEF +5\n", chosen.Name)
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}

	case 2:
		if player.Equip.Weapon.Name == "" {
			fmt.Println("❌ Vous n'avez aucune arme forgée à améliorer.")
			return
		}
		if player.Wallet.Spend(25) {
			player.Equip.Weapon.Atk += 5
			// Met à jour l'arme dans l'inventaire
			ensureWeaponsMap(player)
			player.Weapons[player.Equip.Weapon.Name] = player.Equip.Weapon
			player.RecalculateStatsFromEquipment()
			fmt.Printf("✅ Votre arme %s a été améliorée ! ATK +5\n", player.Equip.Weapon.Name)
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}

	case 0:
		return
	default:
		fmt.Println("❌ Choix invalide, essaie encore.")
	}
}
