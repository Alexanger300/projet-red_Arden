package poisonpot

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Arden/character"
)

// ApplyPoison applique un effet de poison sur la cible pendant X tours
func ApplyPoison(user *character.Character, target *character.Character) {
	if target == nil || !target.IsAlive() {
		fmt.Println("❌ Cible invalide.")
		return
	}

	// Dégâts fixes par tour
	poisonDamage := 5
	poisonDuration := 3 // nombre de tours

	fmt.Printf("☠️ %s utilise une Potion de poison sur %s ! (%d tours)\n", user.Name, target.Name, poisonDuration)

	// On marque la cible comme empoisonnée
	target.ApplyStatus("Poison", poisonDuration, poisonDamage)
}
