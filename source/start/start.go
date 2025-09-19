package start

import (
	"fmt"
	"os"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/inn"
	"github.com/Alexanger300/projet-red_Forge/source/introduction"
	"github.com/Alexanger300/projet-red_Forge/source/save"
)

// petit helper pour choisir un slot (1..3)
func pickSlot(prompt string) int {
	var slot int
	for {
		fmt.Print(prompt)
		fmt.Scan(&slot)
		if slot >= 1 && slot <= 3 {
			return slot
		}
		fmt.Println("❌ Slot invalide. Choisissez 1, 2 ou 3.")
	}
}

func DisplayMenu() {
	for {
		fmt.Println(css.Bold + css.Underline + "=== 🌌 Les Chroniques d'Arden ===" + css.Reset)
		fmt.Println("1 - Nouvelle Partie")
		fmt.Println("2 - Charger une Partie")
		fmt.Println("3 - Qui sont-ils")
		fmt.Println("4 - Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Nouvelle partie --> demander le slot
			slot := pickSlot("Choisissez un slot (1-3) pour la nouvelle partie : ")

			// Si le slot existe déjà, demander confirmation d’écraser
			if save.SlotExists(slot) {
				var ans string
				fmt.Printf("⚠️ Le slot %d contient déjà une sauvegarde. L'écraser ? (o/n) : ", slot)
				fmt.Scan(&ans)
				if ans != "o" && ans != "O" && ans != "oui" && ans != "Oui" {
					fmt.Println("↩️ Retour au menu.")
					break
				}
			}

			// Création du personnage (intro inchangée)
			introduction.Introduction() // ta fonction actuelle crée le perso et lance l’auberge

		case 2:
			// Charger une partie → demander le slot
			slot := pickSlot("Choisissez le slot à charger (1-3) : ")

			state, err := save.LoadGame(slot)
			if err != nil {
				fmt.Println(err) // message déjà clair côté save.LoadGame
				break
			}

			// reconstruire le joueur avec ta fonction existante
			player := character.LoadFromSave(state)

			// lancer l’auberge
			inn.Inn(&player)

		case 3:
			fmt.Println(" ABBA && Steven Spielberg")

		case 4:
			fmt.Println("👋 Au revoir !")
			os.Exit(0)

		default:
			fmt.Println("❌ Choix invalide")
		}
	}
}
