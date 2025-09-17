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
		fmt.Printf("Nom: %s | Classe: %s | Niveau: %d | XP: %d/%d\n",
			player.Name, player.Class, player.Level, player.Exp, player.ExpNext)
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
			player.DisplayInventoryAndEquipment()
		case 5:
			player.HP = player.MaxHP
			player.Mana = player.MaxMana
			fmt.Println("✨ Vous vous reposez et récupérez toute votre énergie !")
		case 6:
			save.SaveGame(save.GameState{
				Name:      player.Name,
				Class:     player.Class,
				Money:     player.Wallet.Amount,
				Progress:  "auberge",
				Inventory: map[string]int{},
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
