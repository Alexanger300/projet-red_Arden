package blacksmith

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/character"
	"github.com/Alexanger300/projet-red_Forge/equipment"
)

// 🔹 Vérifie si le joueur a toutes les ressources demandées
func hasResources(inv map[string]int, cost map[string]int) bool {
	for item, qty := range cost {
		if inv[item] < qty {
			return false
		}
	}
	return true
}

// 🔹 Retire les ressources de l’inventaire
func removeResources(inv map[string]int, cost map[string]int) {
	for item, qty := range cost {
		inv[item] -= qty
		if inv[item] <= 0 {
			delete(inv, item)
		}
	}
}

func Welcome(player *character.Character) {
	fmt.Println("=== Bienvenue chez le forgeron ===")
	fmt.Println("Je peux forger ou améliorer votre équipement.")

	for {
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Forger un équipement")
		fmt.Println("2. Améliorer un équipement")
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
			fmt.Println("❌ Choix invalide.")
		}
	}
}

func forgeEquipment(player *character.Character) {
	classPool, ok := equipment.EquipmentPools[player.Class]
	if !ok {
		fmt.Println("❌ Aucun équipement disponible pour cette classe.")
		return
	}

	fmt.Println("\n--- Forge d'Équipement ---")
	for i, item := range classPool {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(classPool) {
		fmt.Println("❌ Choix invalide.")
		return
	}

	item := classPool[choice-1]
	costGold := 20

	// 🔹 Exemple de ressources nécessaires pour chaque type d’équipement
	// Tu pourras adapter ça comme tu veux selon ton univers
	resourceCosts := map[string]map[string]int{
		"Casque de paladin":    {"Lingot de fer": 2, "Peau de gobelin": 1},
		"Armure bénie":         {"Lingot de fer": 3, "Cristal magique": 1},
		"Jambières lourdes":    {"Lingot de fer": 2, "Cuir de sanglier": 2},
		"Heaume massif":        {"Pierre de vie": 1, "Lingot de fer": 2},
		"Plastron de colosse":  {"Lingot de fer": 4, "Cuir de sanglier": 2},
		"Jambières de pierre":  {"Lingot de fer": 2, "Pierre de vie": 2},
		"Chapeau mystique":     {"Cristal magique": 2, "Parchemin ancien": 1},
		"Robe enchantée":       {"Cristal magique": 3, "Branche d'arbre": 2},
		"Bottes de lévitation": {"Cristal magique": 1, "Fourrure de loup": 2},
		"Capuche de prêtre":    {"Branche d'arbre": 2, "Pierre de vie": 1},
		"Robe de lumière":      {"Branche d'arbre": 3, "Cristal magique": 1},
		"Sandales bénies":      {"Branche d'arbre": 1, "Peau de gobelin": 1},
	}

	required, exists := resourceCosts[item.Name]
	if !exists {
		fmt.Println("❌ Aucune recette trouvée pour cet équipement.")
		return
	}

	// Vérif ressources + or
	if player.Wallet.Amount < costGold {
		fmt.Println("❌ Pas assez d'or !")
		return
	}
	if !hasResources(player.Inventory, required) {
		fmt.Println("❌ Pas assez de ressources !")
		return
	}

	// Consommer ressources + or
	player.Wallet.Amount -= costGold
	removeResources(player.Inventory, required)

	// Équiper l’objet
	player.Equip.Equip(item, player.Class)
	player.RecalculateStatsFromEquipment()
	fmt.Printf("✅ %s forgé et équipé !\n", item.Name)
}

func improveEquipment(player *character.Character) {
	fmt.Println("\n--- Amélioration d'Équipement ---")
	player.Equip.Display()

	fmt.Println("\nQue voulez-vous améliorer ?")
	fmt.Println("1. Tête (+5 DEF, 30 or)")
	fmt.Println("2. Corps (+10 DEF, 50 or)")
	fmt.Println("3. Jambes (+5 DEF, 30 or)")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if player.Wallet.Amount >= 30 && player.Equip.Head.Name != "" {
			player.Wallet.Amount -= 30
			player.Equip.Head.Def += 5
			player.RecalculateStatsFromEquipment()
			fmt.Println("✅ Casque amélioré ! DEF +5")
		} else {
			fmt.Println("❌ Impossible (pas d'or ou pas d'équipement).")
		}
	case 2:
		if player.Wallet.Amount >= 50 && player.Equip.Body.Name != "" {
			player.Wallet.Amount -= 50
			player.Equip.Body.Def += 10
			player.RecalculateStatsFromEquipment()
			fmt.Println("✅ Armure améliorée ! DEF +10")
		} else {
			fmt.Println("❌ Impossible (pas d'or ou pas d'équipement).")
		}
	case 3:
		if player.Wallet.Amount >= 30 && player.Equip.Legs.Name != "" {
			player.Wallet.Amount -= 30
			player.Equip.Legs.Def += 5
			player.RecalculateStatsFromEquipment()
			fmt.Println("✅ Jambières améliorées ! DEF +5")
		} else {
			fmt.Println("❌ Impossible (pas d'or ou pas d'équipement).")
		}
	case 0:
		return
	default:
		fmt.Println("❌ Choix invalide.")
	}
}
