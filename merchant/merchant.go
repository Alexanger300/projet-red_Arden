package merchant

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Arden/character"
	"github.com/Alexanger300/projet-red_Arden/inventory"
)

func Welcome(player *character.Character) {
	fmt.Println("=== 🏪 Bienvenue chez le marchand ===")
	fmt.Println("J'ai toutes sortes d'objets à vendre !")

	for {
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Acheter des objets")
		fmt.Println("2. Vendre des objets")
		fmt.Println("3. Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			buyItems(player)
		case 2:
			sellItems(player)
		case 3:
			fmt.Println("Merci de votre visite, revenez bientôt !")
			return
		default:
			fmt.Println("Choix invalide, essayez encore.")
		}
	}
}

// === Achat d’objets ===
func buyItems(player *character.Character) {
	fmt.Println("\n--- 🛒 Boutique ---")
	fmt.Println("1. Potion de soin (20 or)")
	fmt.Println("2. Élixir de mana (15 or)")
	fmt.Println("3. Cuir de sanglier (10 or)")
	fmt.Println("4. Sac amélioré (+10 slots) (50 or)")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if player.Wallet.Spend(20) {
			inventory.AddItem("Potion de soin", 1)
			fmt.Println("✅ Vous avez acheté une Potion de soin !")
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}
	case 2:
		if player.Wallet.Spend(15) {
			inventory.AddItem("Élixir de mana", 1)
			fmt.Println("✅ Vous avez acheté un Élixir de mana !")
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}
	case 3:
		if player.Wallet.Spend(10) {
			inventory.AddItem("Cuir de sanglier", 1)
			fmt.Println("✅ Vous avez acheté un Cuir de sanglier !")
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}
	case 4:
		if player.Wallet.Spend(50) {
			inventory.UpgradeBag(10) // ⚡ Augmente la capacité max de +10
			fmt.Println("👜 Vous avez acheté un sac amélioré ! Votre inventaire peut contenir plus d'objets.")
		} else {
			fmt.Println("❌ Pas assez d'or.")
		}
	case 0:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}

// === Vente d’objets ===
func sellItems(player *character.Character) {
	fmt.Println("\n--- 💰 Vente ---")
	fmt.Println("1. Vendre une Potion de soin (+10 or)")
	fmt.Println("2. Vendre un Élixir de mana (+7 or)")
	fmt.Println("3. Vendre un Cuir de sanglier (+5 or)")
	fmt.Println("0. Retour")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if inventory.HasItem("Potion de soin", 1) {
			inventory.RemoveItem("Potion de soin", 1)
			player.Wallet.Add(10)
			fmt.Println("✅ Vous avez vendu une Potion de soin.")
		} else {
			fmt.Println("❌ Vous n'avez pas de Potion de soin à vendre.")
		}
	case 2:
		if inventory.HasItem("Élixir de mana", 1) {
			inventory.RemoveItem("Élixir de mana", 1)
			player.Wallet.Add(7)
			fmt.Println("✅ Vous avez vendu un Élixir de mana.")
		} else {
			fmt.Println("❌ Vous n'avez pas d'Élixir de mana à vendre.")
		}
	case 3:
		if inventory.HasItem("Cuir de sanglier", 1) {
			inventory.RemoveItem("Cuir de sanglier", 1)
			player.Wallet.Add(5)
			fmt.Println("✅ Vous avez vendu un Cuir de sanglier.")
		} else {
			fmt.Println("❌ Vous n'avez pas de Cuir de sanglier à vendre.")
		}
	case 0:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
