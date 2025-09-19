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

func TakePotMana(player *character.Character) {
	// Vérifie si le joueur a une potion
	if !inventory.HasItem("Élixir de mana", 1) {
		fmt.Println("❌ Vous n'avez pas d'élixir de mana dans l'inventaire.")
		return
	}
	// Supprime l'élixir de l’inventaire
	inventory.RemoveItem("Élixir de mana", 1)
	// Restaure le mana du joueur
	manaAmount := 20
	player.Mana += manaAmount
	if player.Mana > player.MaxMana {
		player.Mana = player.MaxMana
	}

	// Feedback joueur
	fmt.Printf("✨ %s utilise un Élixir de mana ! (+%d PM)\n", player.Name, manaAmount)
	fmt.Printf("💧 PM actuels : %d/%d\n", player.Mana, player.MaxMana)
}
