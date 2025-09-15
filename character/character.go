package character

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Arden/class"
)

type Character struct {
	Name    string
	Class   string
	HPMax   int
	HP      int
	ManaMax int
	Mana    int
	Atk     int
	Def     int
	Spd     int
	Crit    int
	Weapon  string
}

// Création du personnage avec choix du joueur
func InitCharacter() Character {
	var c Character
	var choiceNumber int
	var confirm string
	confirmed := false

	// Nom
	fmt.Print("Entrez le nom du personnage : ")
	_, err := fmt.Scan(&c.Name)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	fmt.Printf("Voici le nom de votre personnage : %s\n", c.Name)

	// Classe
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
		fmt.Printf("HP: %d | ATK: %d | DEF: %d | Mana: %d | SPD: %d | CRIT: %d%% | Arme: %s\n",
			stats.HPMax, stats.Atk, stats.Def, stats.ManaMax, stats.Spd, stats.Crit, stats.Weapon)

		fmt.Print("Confirmez-vous votre choix ? (Oui/Non) : ")
		fmt.Scan(&confirm)

		if confirm == "Oui" || confirm == "oui" {
			c.Class = className
			c.HPMax = stats.HPMax
			c.HP = stats.HPMax
			c.ManaMax = stats.Mana
			c.Mana = stats.ManaMax
			c.Atk = stats.Atk
			c.Def = stats.Def
			c.Spd = stats.Spd
			c.Crit = stats.Crit
			c.Weapon = stats.Weapon
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
