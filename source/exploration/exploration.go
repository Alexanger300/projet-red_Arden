package exploration

import (
	"fmt"
	"time"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/fight"
	"github.com/Alexanger300/projet-red_Forge/source/monster"
)

// Définition des ressources lootables par zone
var zoneLoots = map[string][]string{
	"Forêt sombre": {
		"Branche d'arbre",
		"Parchemin ancien",
	},
	"Montagnes glacées": {
		"Lingot de fer",
	},
	"Ruines maudites": {
		"Cristal magique",
		"Pierre de vie",
	},
}

var alreadyUsed = false

func Start(player *character.Character) {
	css.Clear()
	text1 := `Les terres d'Arden ne sont plus qu'un champ de ruines.`
	text2 := `Trois routes mènent aux vestiges de l’ancien royaume`
	text3 := ` : les forêts sombres, `
	text4 := `les montagnes glacées,`
	text5 := `et les ruines maudites.`
	text6 := `Chaque pas est un pari contre la mort.`
	if !alreadyUsed {
		for _, c := range text1 {
			fmt.Print(string(c))
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(1 * time.Second)
		for _, c := range text2 {
			fmt.Print(css.Bold + css.Red + string(c) + css.Reset)
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(1 * time.Second)
		for _, c := range text3 {
			fmt.Print(css.Green + string(c) + css.Reset)
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(200 * time.Millisecond)
		for _, c := range text4 {
			fmt.Print(css.Blue + string(c) + css.Reset)
			time.Sleep(50 * time.Millisecond)
		}
		time.Sleep(200 * time.Millisecond)
		for _, c := range text5 {
			fmt.Print(css.Gray + string(c) + css.Reset)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		for _, c := range text6 {
			fmt.Print(css.Bold + css.Red + string(c) + css.Reset)
			time.Sleep(50 * time.Millisecond)
		}
	}
	alreadyUsed = true
	time.Sleep(1 * time.Second)
	for {
		player.DisplayStatsBar()

		fmt.Println("\n=== 🌌 Exploration des terres d'Arden ===")
		fmt.Println(css.Green + "1. Forêt sombre" + css.Reset)
		fmt.Println(css.Blue + "2. Montagnes glacées" + css.Reset)
		fmt.Println(css.Gray + "3. Ruines maudites" + css.Reset)
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

		// Pas besoin de redonner le loot du monstre ici,
		// fight.StartFight s'en occupe déjà

		// Fouille de la zone
		var choix string
		fmt.Printf("\nVoulez-vous fouiller la zone %s pour trouver des ressources ? (oui/non) : ", zoneName)
		fmt.Scan(&choix)

		if choix == "oui" || choix == "Oui" {
			if loots, ok := zoneLoots[zoneName]; ok {
				for _, loot := range loots {
					player.AddItem(loot, 1) // 1 exemplaire de chaque ressource trouvée
					fmt.Printf("Vous trouvez : %s\n", loot)
				}
			}
		} else {
			fmt.Println("👉 Vous retournez directement au village.")
		}

	}
}
