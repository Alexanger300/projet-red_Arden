package start

import (
	"fmt"
	"os"

	"github.com/Alexanger300/projet-red_Forge/character"
	"github.com/Alexanger300/projet-red_Forge/inn"
	"github.com/Alexanger300/projet-red_Forge/introduction"
	"github.com/Alexanger300/projet-red_Forge/money"
	"github.com/Alexanger300/projet-red_Forge/save"
)

// DisplayMenu lance le menu principal du jeu
func DisplayMenu() {
	for {
		fmt.Println("=== 🌌 Les Chroniques d'Arden ===")
		fmt.Println("1 - Nouvelle Partie")
		fmt.Println("2 - Charger une Partie")
		fmt.Println("3 - Options")
		fmt.Println("4 - Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// 🚀 Nouvelle partie
			introduction.Introduction()

		case 2:
			// 📂 Charger sauvegarde
			state := save.LoadGame()
			if state.Name == "" {
				fmt.Println("⚠️ Aucune sauvegarde disponible.")
			} else {
				fmt.Printf("📂 Partie chargée : %s (%s), Or : %d Gold\n",
					state.Name, state.Class, state.Money)

				// 🔹 Reconstruire le joueur
				player := character.Character{
					Name:   state.Name,
					Class:  state.Class,
					Wallet: money.NewGold(state.Money),
				}

				// Relancer depuis l’auberge
				inn.Inn(&player)
			}

		case 3:
			fmt.Println("⚙️ Options (pas encore implémentées)")

		case 4:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)

		default:
			fmt.Println("❌ Choix invalide")
		}
	}
}
