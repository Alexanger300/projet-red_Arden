package playerturn

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/character"
	"github.com/Alexanger300/projet-red_Forge/monster"
)

func PlayerTurnAction(player *character.Character, enemy *monster.Monster) {
	fmt.Printf("\n👉 Tour de %s\n", player.Name)
	fmt.Println("1. Attaque basique")
	fmt.Println("2. Utiliser une compétence")
	fmt.Println("3. Inventaire")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Attaque basique
		damage := player.Atk - enemy.Def
		if damage < 0 {
			damage = 0
		}
		enemy.HP -= damage
		if enemy.HP < 0 {
			enemy.HP = 0
		}
		fmt.Printf("⚔️ %s attaque %s → %d dégâts (HP %d/%d)\n",
			player.Name, enemy.Name, damage, enemy.HP, enemy.HPMax)

	case 2:
		// Compétences
		if len(player.Skills) == 0 {
			fmt.Println("❌ Vous n'avez aucune compétence.")
			return
		}
		fmt.Println("\n--- Compétences ---")
		for i, s := range player.Skills {
			fmt.Printf("%d. %s (Mana: %d)\n", i+1, s.Name, s.ManaCost)
		}
		var idx int
		fmt.Print("Choix : ")
		fmt.Scan(&idx)
		if idx < 1 || idx > len(player.Skills) {
			fmt.Println("❌ Choix invalide.")
			return
		}
		s := player.Skills[idx-1]

		// Vérif mana
		if player.Mana < s.ManaCost {
			fmt.Println("❌ Pas assez de mana.")
			return
		}
		player.Mana -= s.ManaCost

		// Soin (sur soi)
		if s.IsHeal {
			heal := s.Power + (player.Mana / 10)
			player.HP += heal
			if player.HP > player.MaxHP {
				player.HP = player.MaxHP
			}
			fmt.Printf("✨ %s lance %s → +%d HP (HP %d/%d)\n",
				player.Name, s.Name, heal, player.HP, player.MaxHP)
			return
		}

		// Dégâts sur le monstre
		raw := 0
		if s.IsMagic {
			raw = s.Power + (player.Mana / 5)
		} else {
			raw = s.Power + (player.Atk / 2)
		}
		final := raw - enemy.Def
		if final < 0 {
			final = 0
		}
		enemy.HP -= final
		if enemy.HP < 0 {
			enemy.HP = 0
		}
		fmt.Printf("🔥 %s lance %s → %s subit %d dégâts (HP %d/%d)\n",
			player.Name, s.Name, enemy.Name, final, enemy.HP, enemy.HPMax)

	case 3:
		// Inventaire (potions principales)
		if len(player.Inventory) == 0 {
			fmt.Println("🎒 Inventaire vide.")
			return
		}
		items := make([]string, 0, len(player.Inventory))
		fmt.Println("\n--- Inventaire ---")
		i := 1
		for name, qty := range player.Inventory {
			fmt.Printf("%d. %s x%d\n", i, name, qty)
			items = append(items, name)
			i++
		}
		var pick int
		fmt.Print("Choix : ")
		fmt.Scan(&pick)
		if pick < 1 || pick > len(items) {
			fmt.Println("❌ Choix invalide.")
			return
		}
		item := items[pick-1]

		switch item {
		case "Potion de soin":
			if player.RemoveItem(item, 1) {
				heal := 50
				player.HP += heal
				if player.HP > player.MaxHP {
					player.HP = player.MaxHP
				}
				fmt.Printf("🍷 %s boit une potion → +%d HP (HP %d/%d)\n",
					player.Name, heal, player.HP, player.MaxHP)
			} else {
				fmt.Println("❌ Plus de potions de soin.")
			}

		case "Potion de poison":
			if player.RemoveItem(item, 1) {
				enemy.ApplyStatus("Poison", 3, 5)
			} else {
				fmt.Println("❌ Plus de potions de poison.")
			}

		default:
			fmt.Println("❌ Objet non utilisable en combat.")
		}

	default:
		fmt.Println("❌ Action invalide. Tour perdu.")
	}
}
