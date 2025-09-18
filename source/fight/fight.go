package fight

import (
	"fmt"
	"time"

	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/monster"
)

// Fonction principale de combat
func StartFight(player *character.Character, enemy *monster.Monster) bool {
	fmt.Printf("\n⚔️ Un %s apparaît ! (%d/%d HP)\n", enemy.Name, enemy.HP, enemy.HPMax)

	// Boucle de combat
	for player.IsAlive() && enemy.IsAlive() {
		// Tour du joueur
		escaped := playerTurn(player, enemy)
		if escaped {
			fmt.Println("Vous avez réussi à fuir le combat ! Cependant, vous avez perdu 10 d'or en fuyant.")
			return false
		}

		if !enemy.IsAlive() {
			break
		}

		// Mise à jour statuts du monstre
		enemy.UpdateStatuses()
		if !enemy.IsAlive() {
			break
		}

		// --- Tour du monstre ---
		monsterTurn(player, enemy)
		if !player.IsAlive() {
			break
		}

		// Mise à jour statuts du joueur
		player.UpdateStatuses()
	}

	// Fin du combat
	if player.IsAlive() {
		fmt.Printf("\n🏆 Victoire ! Vous avez vaincu %s !\n", enemy.Name)

		// 🔹 Ajout du loot du monstre (si défini)
		if enemy.Loot != "" {
			player.AddItem(enemy.Loot, 1) // 1 exemplaire du loot
			fmt.Printf("📦 Vous récupérez : %s\n", enemy.Loot)
		}

		return true
	}

	fmt.Println("\n💀 Vous êtes mort...")
	return false
}

// Tour du joueur
func playerTurn(player *character.Character, enemy *monster.Monster) bool {
	for {
		player.DisplayStatsBar()

		fmt.Printf("\n🎯 Tour de %s (HP %d/%d | Mana %d/%d)\n",
			player.Name, player.HP, player.MaxHP, player.Mana, player.MaxMana)
		fmt.Println("1. Attaque basique ⚔️")
		fmt.Println("2. Utiliser une compétence 🔮")
		fmt.Println("3. Inventaire 🎒")
		fmt.Println("4. Fuir 🏃‍♂️")

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
			time.Sleep(1 * time.Second)
			return false // tour consommé, pas de fuite

		case 2:
			// Sous-menu compétences
			if len(player.Skills) == 0 {
				fmt.Println("❌ Vous n'avez aucune compétence.")
				continue
			}

			for {
				fmt.Println("\n--- Compétences ---")
				for i, s := range player.Skills {
					fmt.Printf("%d. %s (Mana %d) → %s\n", i+1, s.Name, s.ManaCost, s.Description)
				}
				fmt.Println("0. Retour")

				var idx int
				fmt.Print("Votre choix : ")
				fmt.Scan(&idx)

				if idx == 0 {
					break
				}
				if idx < 1 || idx > len(player.Skills) {
					fmt.Println("❌ Choix invalide.")
					continue
				}

				player.UseSkillOnMonster(player.Skills[idx-1].Name, enemy)
				time.Sleep(1 * time.Second)
				return false // tour consommé, pas de fuite
			}
			continue

		case 3:
			// Sous-menu inventaire
			if len(player.Inventory) == 0 {
				fmt.Println("❌ Inventaire vide.")
				continue
			}

			for {
				fmt.Println("\n--- Inventaire ---")
				i := 1
				items := []string{}
				for item, qty := range player.Inventory {
					if qty > 0 {
						fmt.Printf("%d. %s x%d\n", i, item, qty)
						items = append(items, item)
						i++
					}
				}
				if len(items) == 0 {
					fmt.Println("❌ Inventaire vide.")
					break
				}
				fmt.Println("0. Retour")

				var idx int
				fmt.Print("Votre choix : ")
				fmt.Scan(&idx)

				if idx == 0 {
					break
				}
				if idx < 1 || idx > len(items) {
					fmt.Println("❌ Choix invalide.")
					continue
				}

				// Vérifie l’objet choisi
				chosen := items[idx-1]
				if chosen == "Potion de soin" {
					player.UseItem("Potion de soin", player) // soigne le joueur
				} else {
					player.UseItemOnMonster(chosen, enemy) // ex: potion de poison
				}

				time.Sleep(1 * time.Second)
				return false // tour consommé, pas de fuite
			}
			continue

		case 4:
			// Fuite
			player.Wallet.Spend(10) // coût de la fuite
			return true
		}
		return false
	}
}

// Tour du monstre
func monsterTurn(player *character.Character, enemy *monster.Monster) {
	fmt.Printf("\n👹 Tour de %s\n", enemy.Name)

	if enemy.HP < enemy.HPMax/2 {
		fmt.Printf("%s utilise son attaque spéciale : %s\n", enemy.Name, enemy.SpecialAtk)
		damage := (enemy.Atk*2 - (player.Def / 2))
		if damage < 1 {
			damage = 1
		}
		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Printf("%s inflige %d dégâts à %s (HP: %d/%d)\n",
			enemy.Name, damage, player.Name, player.HP, player.MaxHP)
	} else {
		fmt.Printf("%s utilise son attaque basique : %s\n", enemy.Name, enemy.BasicAtk)
		damage := enemy.Atk - (player.Def / 2)
		if damage < 1 {
			damage = 1
		}
		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Printf("%s inflige %d dégâts à %s (HP: %d/%d)\n",
			enemy.Name, damage, player.Name, player.HP, player.MaxHP)
	}

	time.Sleep(1 * time.Second)
}
