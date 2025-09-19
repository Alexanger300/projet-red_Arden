package monsterturn

import (
	"fmt"
	"math/rand"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/monster"
)

func EnemyTurn(enemy *monster.Monster, player *character.Character) {
	fmt.Printf("\n👹 Tour de %s\n", css.Red+enemy.Name+css.Reset)

	// 0 → basique, 1 → spéciale
	if rand.Intn(3) < 2 {
		// Attaque basique
		dmg := enemy.Atk - player.Def
		if dmg < 0 {
			dmg = 0
		}
		player.HP -= dmg
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Printf("%s utilise %s → %s subit %d dégâts (HP %d/%d)\n",
			enemy.Name, enemy.BasicAtk, player.Name, dmg, player.HP, player.MaxHP)
	} else {
		// Spéciale
		dmg := (enemy.Atk * 2) - player.Def
		if dmg < 0 {
			dmg = 0
		}
		player.HP -= dmg
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Printf("%s utilise %s → %s subit %d dégâts (HP %d/%d)\n",
			enemy.Name, enemy.SpecialAtk, player.Name, dmg, player.HP, player.MaxHP)
	}
}
