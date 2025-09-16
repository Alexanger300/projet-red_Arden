package poisonpot

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/character"
)

// Constants pour équilibrer la potion de poison
const (
	poisonDuration = 3 // dure 3 tours
	poisonDamage   = 5 // 5 dégâts par tour
)

// UsePoisonPotion applique un poison à une cible
func UsePoisonPotion(user *character.Character, target *character.Character) {
	// Vérifier si le joueur a bien une potion (facultatif : lier à inventory)
	fmt.Printf("🧪 %s utilise une potion de poison sur %s !\n", user.Name, target.Name)

	// Appliquer le statut Poison
	target.ApplyStatus("Poison", poisonDuration, poisonDamage)
}
