package playerturn

import (
	"fmt"
	"time"

	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/monster"
)

func playerTurn(player *character.Character, enemy *monster.Monster) {
	fmt.Printf("\n🎯 Tour de %s (HP %d/%d | Mana %d/%d)\n",
		player.Name, player.HP, player.MaxHP, player.Mana, player.MaxMana)
	fmt.Println("1. Attaque basique ⚔️")
	fmt.Println("2. Utiliser une compétence 🔮")
	fmt.Println("3. Inventaire 🎒")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Attaque basique
		damage := player.Atk - (enemy.Def / 2)
		if damage < 1 {
			damage = 1
		}
		enemy.HP -= damage
		if enemy.HP < 0 {
			enemy.HP = 0
		}
		fmt.Printf("⚔️ %s attaque %s → %d dégâts (HP ennemi : %d/%d)\n",
			player.Name, enemy.Name, damage, enemy.HP, enemy.HPMax)

	case 2:
		if len(player.Skills) == 0 {
			fmt.Println("❌ Vous n'avez aucune compétence.")
			return // pas de skill dispo → on quitte proprement sans perdre le combat
		}

		fmt.Println("\n--- Compétences ---")
		for i, s := range player.Skills {
			fmt.Printf("%d. %s (Mana %d) → %s\n", i+1, s.Name, s.ManaCost, s.Description)
		}
		fmt.Println("0. Retour")

		var idx int
		fmt.Print("Votre choix : ")
		fmt.Scan(&idx)

		if idx == 0 {
			fmt.Println("↩️ Retour au menu principal.")
			playerTurn(player, enemy) // 🔁 relance le menu principal sans perdre le tour
			return
		}
		if idx < 1 || idx > len(player.Skills) {
			fmt.Println("❌ Choix invalide.")
			playerTurn(player, enemy) // 🔁 relance le menu principal sans perdre le tour
			return
		}

		player.UseSkillOnMonster(player.Skills[idx-1].Name, enemy)

	case 3:
		if len(player.Inventory) == 0 {
			fmt.Println("❌ Inventaire vide.")
			return
		}

		fmt.Println("\n--- Inventaire ---")
		i := 1
		items := []string{}
		for item, qty := range player.Inventory {
			fmt.Printf("%d. %s x%d\n", i, item, qty)
			items = append(items, item)
			i++
		}
		fmt.Println("0. Retour")

		var idx int
		fmt.Print("Votre choix : ")
		fmt.Scan(&idx)

		if idx == 0 {
			fmt.Println("↩️ Retour au menu principal.")
			playerTurn(player, enemy) // 🔁 relance le menu principal sans perdre le tour
			return
		}
		if idx < 1 || idx > len(items) {
			fmt.Println("❌ Choix invalide.")
			playerTurn(player, enemy) // 🔁 relance le menu principal sans perdre le tour
			return
		}

		player.UseItemOnMonster(items[idx-1], enemy)

	default:
		fmt.Println("❌ Choix invalide.")
		playerTurn(player, enemy) // 🔁 relance le menu principal sans perdre le tour
		return
	}

	time.Sleep(1 * time.Second)
}
