package exploration

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/character"
	"github.com/Alexanger300/projet-red_Forge/inventory"
)

func Start(player *character.Character) {
	fmt.Println("\n=== 🌌 Exploration des terres d'Arden ===")
	fmt.Println("1. 🌲 Forêt sombre")
	fmt.Println("2. 🏔️ Montagnes glacées")
	fmt.Println("3. ☠️ Ruines maudites")
	fmt.Println("0. Retour")

	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		forest(player)
	case 2:
		mountains(player)
	case 3:
		ruins(player)
	case 0:
		return // on ne va pas vers auberge ici, juste on sort
	default:
		fmt.Println("❌ Choix invalide")
	}
}

func forest(player *character.Character) {
	fmt.Println("🌲 Gobelin surgit ! (combat placeholder)")
	inventory.AddItem("Peau de gobelin", 1)
}

func mountains(player *character.Character) {
	fmt.Println("🏔️ Loup affamé ! (combat placeholder)")
	inventory.AddItem("Fourrure de loup", 1)
}

func ruins(player *character.Character) {
	fmt.Println("☠️ Sanglier monstrueux ! (combat placeholder)")
	inventory.AddItem("Peau de sanglier", 1)
}
