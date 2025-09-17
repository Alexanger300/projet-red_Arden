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
		//  Tour du joueur
		playerTurn(player, enemy)
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
		return true
	}
	fmt.Println("\n💀 Vous êtes mort...")
	return false
}

//	Tour du joueur
//
// Ne consomme le tour QUE si une action valide est effectuée.
// - 1: Attaque → termine le tour
// - 2: Compétences → sous-menu avec "0. Retour". Retour/choix invalide ne consomment PAS le tour
// - 3: Inventaire  → sous-menu avec "0. Retour". Retour/choix invalide ne consomment PAS le tour
func playerTurn(player *character.Character, enemy *monster.Monster) {
	for {
		player.DisplayStatsBar()

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
			// Attaque basique (formule équilibrée déjà utilisée chez toi)
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
			return // ✅ tour consommé

		case 2:
			// Sous-menu compétences (avec retour)
			if len(player.Skills) == 0 {
				fmt.Println("❌ Vous n'avez aucune compétence.")
				continue // ❌ pas de compétence → on ne consomme pas le tour
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
					// ↩️ retour au menu principal du tour sans perdre le tour
					break
				}
				if idx < 1 || idx > len(player.Skills) {
					fmt.Println("❌ Choix invalide.")
					continue // re-afficher la liste de compétences
				}

				// On lance la compétence choisie
				player.UseSkillOnMonster(player.Skills[idx-1].Name, enemy)
				time.Sleep(1 * time.Second)
				return // ✅ tour consommé (on a effectivement agi)
			}

			// si on a “break” depuis le sous-menu → on repart en haut du for et on repropose 1/2/3
			continue

		case 3:
			// Sous-menu inventaire (avec retour)
			if len(player.Inventory) == 0 {
				fmt.Println("❌ Inventaire vide.")
				continue // pas de tour consommé
			}

			// construire une liste d’items indexés
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
					break // revient au menu principal du tour
				}
				fmt.Println("0. Retour")

				var idx int
				fmt.Print("Votre choix : ")
				fmt.Scan(&idx)

				if idx == 0 {
					// ↩️ retour au menu principal du tour
					break
				}
				if idx < 1 || idx > len(items) {
					fmt.Println("❌ Choix invalide.")
					continue // re-afficher inventaire
				}

				// Utiliser l’objet sélectionné sur le monstre (ex: potion de poison)
				player.UseItemOnMonster(items[idx-1], enemy)
				time.Sleep(1 * time.Second)
				return // ✅ tour consommé
			}

			continue // revient au menu principal du tour

		default:
			fmt.Println("❌ Choix invalide.")
			// on REPROPOSE le menu au lieu de consommer le tour
			continue
		}
	}
}

// === Tour du monstre ===
func monsterTurn(player *character.Character, enemy *monster.Monster) {
	fmt.Printf("\n👹 Tour de %s\n", enemy.Name)

	// Attaque spéciale si le monstre est à < 50% HP, sinon basique
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
