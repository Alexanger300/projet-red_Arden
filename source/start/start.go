package start

import (
	"fmt"
	"os"

	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/inn"
	"github.com/Alexanger300/projet-red_Forge/source/introduction"
	"github.com/Alexanger300/projet-red_Forge/source/save"
)

func DisplayMenu() {
	for {
		fmt.Println("=== 🌌 Les Chroniques d'Arden ===")
		fmt.Println("1 - Nouvelle Partie")
		fmt.Println("2 - Charger une Partie")
		fmt.Println("3 - Qui sont-ils ?")
		fmt.Println("4 - Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			introduction.Introduction()

		case 2:
			state := save.LoadGame()
			if state.Name == "" {
				fmt.Println("⚠️ Aucune sauvegarde disponible.")
				break
			}

			fmt.Printf("📂 Partie chargée : %s (%s), Or : %d Gold\n",
				state.Name, state.Class, state.Money)

			// ✅ Reconstruire le personnage depuis la sauvegarde
			player := character.LoadFromSave(state)

			inn.Inn(&player)

		case 3:
			fmt.Println("ABBA / Steven Spielberg")
		case 4:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("❌ Choix invalide")
		}
	}
}
