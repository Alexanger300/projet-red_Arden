package exploration

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/fight"
	"github.com/Alexanger300/projet-red_Forge/source/inventory"
	"github.com/Alexanger300/projet-red_Forge/source/monster"
)

// Définition des ressources lootables par zone
var zoneLoots = map[string][]string{
	"Forêt sombre":      {"Branche d'arbre, parchemin ancien"},
	"Montagnes glacées": {"Lingot de fer"},
	"Ruines maudites":   {"Cristal magique", "Pierre de vie"},
}

func Start(player *character.Character) {
	for {
		player.DisplayStatsBar()

		fmt.Println("\n=== 🌌 Exploration des terres d'Arden ===")
		fmt.Println("1. Forêt sombre")
		fmt.Println("2. Montagnes glacées")
		fmt.Println("3. Ruines maudites")
		fmt.Println("0. Retour")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			exploreZone(player, monster.NewGoblin(), "Forêt sombre")
		case 2:
			exploreZone(player, monster.NewWolf(), "Montagnes glacées")
		case 3:
			exploreZone(player, monster.NewBoar(), "Ruines maudites")
		case 0:
			return
		default:
			fmt.Println("❌ Choix invalide")
		}
	}
}

// exploreZone lance un combat dans une zone donnée
func exploreZone(player *character.Character, enemy *monster.Monster, zoneName string) {
	fmt.Printf("\n⚔️ Vous entrez dans la zone et rencontrez un %s !\n", enemy.Name)

	// Combat tour par tour
	victory := fight.StartFight(player, enemy)

	if victory {
		fmt.Printf("🏆 Vous avez vaincu le %s !\n", enemy.Name)

		// Gain d’XP
		player.GainExp(enemy.ExpReward)

		// Gain d’or
		player.Wallet.Add(enemy.GoldReward)
		fmt.Printf("💰 Vous gagnez %d or. Total : %d\n", enemy.GoldReward, player.Wallet.Amount)

		// Loot de monstre
		if enemy.Loot != "" {
			inventory.AddItem(enemy.Loot, 1)
			fmt.Printf("📦 Vous récupérez : %s\n", enemy.Loot)
		}

		// 🔹 Choix de loot de la zone
		var choix string
		fmt.Printf("\nVoulez-vous fouiller la zone %s pour trouver des ressources ? (oui/non) : ", zoneName)
		fmt.Scan(&choix)

		if choix == "oui" || choix == "Oui" {
			if loots, ok := zoneLoots[zoneName]; ok {
				for _, item := range loots {
					inventory.AddItem(item, 1)
					fmt.Printf(" Vous trouvez : %s\n", item)
				}
			}
		} else {
			fmt.Println("👉 Vous retournez directement au village.")
		}

	} else {
		fmt.Println("💀 Vous avez été vaincu… Retour à l'auberge.")
	}
}
