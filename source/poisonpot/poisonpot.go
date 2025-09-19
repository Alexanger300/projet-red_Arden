package poisonpot

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/source/character"
)

const (
	poisonDuration = 3
	poisonDamage   = 5
)

// Application de l'effet de poison (si le personnage a une potion de poison)
func UsePoisonPotion(user *character.Character, target *character.Character) {
	if user.RemoveItem("Potion de poison", 1) {
		fmt.Printf("🧪 %s utilise une potion de poison sur %s !\n", user.Name, target.Name)
		target.ApplyStatus("Poison", poisonDuration, poisonDamage)
	} else {
		fmt.Printf("❌ %s n’a plus de potion de poison !\n", user.Name)
	}
}
