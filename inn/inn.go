package inn

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/blacksmith"
	"github.com/Alexanger300/projet-red_Forge/character"
	"github.com/Alexanger300/projet-red_Forge/exploration"
	"github.com/Alexanger300/projet-red_Forge/merchant"
	"github.com/Alexanger300/projet-red_Forge/save"
)

// Auberge : lieu où le joueur se réveille et choisit où aller
func Inn(player *character.Character) {
	for {
		fmt.Println("\n=== 🏰 Auberge d'Arden ===")
		fmt.Printf("👤 %s le %s | ❤️ %d/%d PV | 🔮 %d/%d Mana | 💰 %d Gold\n",
			player.Name, player.Class, player.HP, player.MaxHP, player.Mana, player.MaxMana, player.Wallet.Amount)

		fmt.Println("\nVous vous réveillez dans une petite chambre d'auberge.")
		fmt.Println("En sortant, vous voyez la place du village animée.")
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Aller voir le forgeron ⚒️")
		fmt.Println("2. Aller voir le marchand 🏪")
		fmt.Println("3. Explorer les terres d'Arden 🌌")
		fmt.Println("4. Se reposer (récupérer toute votre vie et mana) 🛏️")
		fmt.Println("5. Sauvegarder et quitter 💾")
		fmt.Println("6. Quitter sans sauvegarder 🚪")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			blacksmith.Welcome(player)
		case 2:
			merchant.Welcome(player)
		case 3:
			exploration.Start(player)
		case 4:
			player.HP = player.MaxHP
			player.Mana = player.MaxMana
			fmt.Println("✨ Vous vous reposez et récupérez toute votre énergie !")
		case 5:
			save.SaveGame(save.GameState{
				Name:      player.Name,
				Class:     player.Class,
				Money:     player.Wallet.Amount,
				Progress:  "auberge",
				Inventory: map[string]int{},
			})
			fmt.Println("💾 Partie sauvegardée. À bientôt !")
			return
		case 6:
			fmt.Println("👋 Vous quittez l'auberge sans sauvegarder.")
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}
