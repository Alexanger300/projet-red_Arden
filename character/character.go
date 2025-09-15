package character

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Arden/class"
	"github.com/Alexanger300/projet-red_Arden/money"
)

type Character struct {
	Name   string
	Class  string
	HP     int
	Mana   int
	Atk    int
	Def    int
	Spd    int
	Crit   int
	Weapon string
	Wallet money.Money
}

// Création du personnage avec intro + choix
func InitCharacter() Character {
	var c Character
	var choiceNumber int
	var confirm string
	confirmed := false

	// INTRODUCTION
	fmt.Println("Au cœur d'un royaume déchiré par la guerre, les terres d'Arden sombrent dans les flammes.")
	fmt.Println("Les dragons ravagent les villages, les armées s'entre-déchirent, et la peste ronge les survivants.")
	fmt.Println()
	fmt.Println("Dans ce chaos, quatre figures se lèvent.")
	fmt.Println("Chacune guidée par un destin différent, mais toutes appelées par la même prophétie :")
	fmt.Println()
	fmt.Println("\"Quand les ténèbres couvriront le ciel, un héros renaîtra de la poussière.\"")
	fmt.Println()
	fmt.Println("À toi de choisir ton rôle dans cette épopée...")
	fmt.Println()

	// NOM DU PERSONNAGE
	fmt.Print("Entrez le nom de votre personnage : ")
	_, err := fmt.Scan(&c.Name)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	fmt.Printf("Voici le nom de votre personnage : %s\n", c.Name)

	// CHOIX DE LA CLASSE
	for !confirmed {
		fmt.Println("\nQuelle Classe voulez-vous ?")
		fmt.Println("1: Paladin")
		fmt.Println("2: Géant")
		fmt.Println("3: Mage")
		fmt.Println("4: Guérisseur")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choiceNumber)

		var className string
		switch choiceNumber {
		case 1:
			className = "Paladin"
		case 2:
			className = "Géant"
		case 3:
			className = "Mage"
		case 4:
			className = "Guérisseur"
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		// Récupérer stats + description depuis class
		stats := class.Classes[className]
		fmt.Printf("\n%s → %s\n", className, stats.Description)
		fmt.Printf("PV: %d | ATK: %d | DEF: %d | Mana: %d | SPD: %d | CRIT: %d%% | Arme: %s\n",
			stats.Pv, stats.Atk, stats.Def, stats.Mana, stats.Spd, stats.Crit, stats.Weapon)

		fmt.Print("Confirmez-vous votre choix ? (Oui/Non) : ")
		fmt.Scan(&confirm)

		if confirm == "Oui" || confirm == "oui" {
			c.Class = className
			c.HP = stats.Pv
			c.Mana = stats.Mana
			c.Atk = stats.Atk
			c.Def = stats.Def
			c.Spd = stats.Spd
			c.Crit = stats.Crit
			c.Weapon = stats.Weapon
			c.Wallet = money.NewGold(100)
			confirmed = true
		}
	}

	return c
}

// Affichage du personnage final
func (c Character) Display() {
	fmt.Println("\n--- Personnage créé avec succès ---")
	fmt.Println("Nom   :", c.Name)
	fmt.Println("Classe:", c.Class)
}
