package character

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Arden/class"
	"github.com/Alexanger300/projet-red_Arden/money"
	"github.com/Alexanger300/projet-red_Arden/skills"
)

type Character struct {
	Name    string
	Class   string
	HP      int
	MaxHP   int
	Mana    int
	MaxMana int
	Atk     int
	Def     int
	Spd     int
	Crit    int
	Weapon  string
	Wallet  money.Money
	Skills  []skills.Skill
}

// Création du personnage
func InitCharacter() Character {
	var c Character
	var choiceNumber int
	var confirm string
	confirmed := false

	// Nom du personnage
	fmt.Print("Entrez le nom du personnage : ")
	_, err := fmt.Scan(&c.Name)
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	fmt.Printf("Voici le nom de votre personnage : %s\n", c.Name)

	// Choix de la classe
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

		// Récupérer les stats depuis "class"
		stats := class.Classes[className]
		fmt.Printf("\n%s → %s\n", className, stats.Description)
		fmt.Printf("PV: %d | ATK: %d | DEF: %d | Mana: %d | SPD: %d | CRIT: %d%% | Arme: %s\n",
			stats.HP, stats.Atk, stats.Def, stats.Mana, stats.Spd, stats.Crit, stats.Weapon)

		fmt.Print("Confirmez-vous votre choix ? (Oui/Non) : ")
		fmt.Scan(&confirm)

		if confirm == "Oui" || confirm == "oui" {
			c.Class = className
			c.HP = stats.HP
			c.MaxHP = stats.HP
			c.Mana = stats.Mana
			c.MaxMana = stats.Mana
			c.Atk = stats.Atk
			c.Def = stats.Def
			c.Spd = stats.Spd
			c.Crit = stats.Crit
			c.Weapon = stats.Weapon
			c.Wallet = money.Money{Amount: 100, Currency: "or"}
			c.Skills = skills.ClassSkills[className]
			confirmed = true
		}
	}

	return c
}

// Vérifie si le joueur est vivant
func (c *Character) IsAlive() bool {
	if c.HP <= 0 {
		fmt.Printf("\n💀 %s est mort... Vous devez recommencer l'aventure.\n", c.Name)
		return false
	}
	return true
}

// Apprendre un nouveau sort
func (c *Character) LearnSkill(newSkill skills.Skill) {
	if newSkill.Name == "Boule de feu" && c.Class != "Mage" {
		fmt.Println("❌ Seul un Mage peut apprendre ce sort.")
		return
	}
	for _, s := range c.Skills {
		if s.Name == newSkill.Name {
			fmt.Println("❌ Vous connaissez déjà ce sort :", newSkill.Name)
			return
		}
	}
	c.Skills = append(c.Skills, newSkill)
	fmt.Println("✨ Nouveau sort appris :", newSkill.Name)
}

// Utiliser un sort sur une cible
func (c *Character) UseSkill(skillName string, target *Character) {
	for _, s := range c.Skills {
		if s.Name == skillName {
			// Vérif restriction Mage
			if s.Name == "Boule de feu" && c.Class != "Mage" {
				fmt.Println("❌ Seul un Mage peut utiliser ce sort.")
				return
			}

			// Vérif mana
			if c.Mana < s.ManaCost {
				fmt.Println("❌ Pas assez de mana pour lancer", s.Name)
				return
			}

			// Consommer le mana
			c.Mana -= s.ManaCost

			// Si c’est un soin
			if s.IsHeal {
				healAmount := s.Power + (c.Mana / 10)
				target.HP += healAmount
				if target.HP > target.MaxHP {
					target.HP = target.MaxHP
				}
				fmt.Printf("✨ %s utilise %s → %s récupère %d PV (HP: %d/%d)\n",
					c.Name, s.Name, target.Name, healAmount, target.HP, target.MaxHP)
				return
			}

			// Sinon → attaque
			var rawDamage int
			if s.IsMagic {
				rawDamage = s.Power + (c.Mana / 5)
			} else {
				rawDamage = s.Power + (c.Atk / 2)
			}

			// Réduction par la DEF
			finalDamage := rawDamage - target.Def
			if finalDamage < 0 {
				finalDamage = 0
			}

			// Appliquer les dégâts
			target.HP -= finalDamage
			if target.HP < 0 {
				target.HP = 0
			}

			fmt.Printf("🔥 %s utilise %s sur %s → %d dégâts (HP restants : %d/%d)\n",
				c.Name, s.Name, target.Name, finalDamage, target.HP, target.MaxHP)

			// Vérifier la mort
			if target.HP == 0 {
				fmt.Printf("💀 %s est mort !\n", target.Name)
			}
			return
		}
	}
	fmt.Println("❌ Sort inconnu :", skillName)
}

// Afficher les compétences
func (c *Character) ShowSkills() {
	if len(c.Skills) == 0 {
		fmt.Println("Aucun sort connu.")
		return
	}
	fmt.Println("\n=== Compétences connues ===")
	for _, s := range c.Skills {
		fmt.Printf("- %s (Mana: %d) → %s\n", s.Name, s.ManaCost, s.Description)
	}
}

// Afficher résumé du personnage
func (c Character) Display() {
	fmt.Println("\n--- Personnage créé avec succès ---")
	fmt.Println("Nom   :", c.Name)
	fmt.Println("Classe:", c.Class)
}
