package takepot

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/inventory"
)

// TakePot : consomme une potion de soin et soigne le personnage
func TakePot(player *character.Character) {
	// Vérifier si le joueur a une potion
	if !inventory.HasItem("Potion de soin", 1) {
		fmt.Println("❌ Vous n'avez pas de potion de soin dans l'inventaire.")
		return
	}

	// Supprimer la potion de l’inventaire
	inventory.RemoveItem("Potion de soin", 1)

	// Soigner le joueur
	healAmount := 20
	player.HP += healAmount
	if player.HP > player.MaxHP {
		player.HP = player.MaxHP
	}

	// Feedback joueur
	fmt.Printf("✨ %s utilise une Potion de soin ! (+%d PV)\n", player.Name, healAmount)
	fmt.Printf("❤️ PV actuels : %d/%d\n", player.HP, player.MaxHP)
}
