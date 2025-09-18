package inn

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/source/blacksmith"
	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/exploration"
	"github.com/Alexanger300/projet-red_Forge/source/merchant"
	"github.com/Alexanger300/projet-red_Forge/source/save"
)

// Auberge : lieu où le joueur se réveille et choisit où aller
func Inn(player *character.Character) {
	for {
		player.DisplayStatsBar()

		fmt.Println("\n=== 🏰 Auberge d'Arden ===")
		fmt.Printf("Nom: %s | Classe: %s | Niveau: %d | XP: %d/%d\n | Argent: %d gold\n",
			player.Name, player.Class, player.Level, player.Exp, player.ExpNext, player.Wallet.Amount)
		fmt.Printf("HP: %d/%d | Mana: %d/%d | ATK: %d | DEF: %d | SPD: %d | CRIT: %d%%\n",
			player.HP, player.MaxHP, player.Mana, player.MaxMana, player.Atk, player.Def, player.Spd, player.Crit)

		fmt.Println("\nVous vous réveillez dans une petite chambre d'auberge.")
		fmt.Println("En sortant, vous voyez la place du village animée.")
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Aller voir le forgeron ⚒️")
		fmt.Println("2. Aller voir le marchand 🏪")
		fmt.Println("3. Explorer les terres d'Arden 🌌")
		fmt.Println("4. Voir l'inventaire et l'équipement 🧳")
		fmt.Println("5. Se reposer (Récupérer toute votre vie et mana) 🛏️")
		fmt.Println("6. Sauvegarder et quitter 💾")
		fmt.Println("7. Quitter sans sauvegarder 🚪")

		var choix int
		fmt.Print("Votre choix : ")
		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Erreur de saisie!", err)
			continue
		}

		switch choix {
		case 1:
			blacksmith.Welcome(player)
		case 2:
			merchant.Welcome(player)
		case 3:
			exploration.Start(player)
		case 4:
			fmt.Println("\n=== 🧳 Inventaire et Équipement ===")
			player.DisplayInventoryAndEquipment()
			fmt.Println("\n0. Retour")
			var subChoice int
			fmt.Print("Votre choix : ")
			_, err := fmt.Scan(&subChoice)
			if err != nil {
				fmt.Println("Entrée invalide!", err)
				continue
			}
			if subChoice != 0 {
				fmt.Println("❌ Choix invalide. Tapez 0 pour revenir.")
			}
		case 5:
			player.HP = player.MaxHP
			player.Mana = player.MaxMana
			fmt.Println("✨ Vous vous reposez et récupérez toute votre énergie !")
		case 6:
			var slot int
			fmt.Print("Choisissez le slot de sauvegarde (1-3) : ")
			fmt.Scan(&slot)

			save.SaveGame(save.GameState{
				SlotID:    slot,
				Name:      player.Name,
				Class:     player.Class,
				Money:     player.Wallet.Amount,
				Progress:  "auberge",
				Inventory: player.Inventory,
			})
			fmt.Println("💾 Partie sauvegardée. À bientôt !")
			return

		case 7:
			fmt.Println("👋 Vous quittez l'auberge sans sauvegarder.")
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}
