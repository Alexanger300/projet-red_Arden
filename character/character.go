package character

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/class"
	"github.com/Alexanger300/projet-red_Forge/equipment"
	"github.com/Alexanger300/projet-red_Forge/money"
	"github.com/Alexanger300/projet-red_Forge/monster"
	"github.com/Alexanger300/projet-red_Forge/save"
	"github.com/Alexanger300/projet-red_Forge/skills"
)

// === Gestion des statuts ===
type Status struct {
	Name     string // Nom du statut (ex: Poison)
	Duration int    // Nombre de tours restants
	Damage   int    // Dégâts par tour (si applicable)
}

// === Structure du personnage ===
type Character struct {
	Name    string
	Gender  string
	Class   string
	Level   int
	Exp     int
	ExpNext int

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

	Statuses  []Status
	Inventory map[string]int         //  Inventaire (objet → quantité)
	Equip     equipment.EquipmentSet //  Ensemble d’équipements
}

// === Création du personnage ===
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

	// Sexe du personnage
	fmt.Print("Choisissez le sexe (Homme/Femme/Autre) : ")
	fmt.Scan(&c.Gender)

	fmt.Printf("Voici le nom de votre personnage : %s (%s)\n", c.Name, c.Gender)

	// Choix de la classe
	for !confirmed {
		fmt.Println("\nQuelle Classe voulez-vous ?")
		fmt.Println("1: Paladin ⚔️")
		fmt.Println("2: Géant 🪓")
		fmt.Println("3: Mage 🔮")
		fmt.Println("4: Guérisseur ✨")
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
			c.Wallet = money.Money{Amount: 100, Currency: "Gold"}
			c.Skills = skills.ClassSkills[className]

			// Expérience
			c.Level = 1
			c.Exp = 0
			c.ExpNext = 10

			// Inventaire de base
			c.Inventory = map[string]int{
				"Potion de soin":   2,
				"Potion de poison": 1,
			}

			// Équipement vide au départ
			c.Equip = equipment.EquipmentSet{}

			confirmed = true
		}
	}

	return c
}

// === Vérifie si le joueur est vivant ===
func (c *Character) IsAlive() bool {
	return c.HP > 0
}

// === Gagner de l’expérience ===
func (c *Character) GainExp(amount int) {
	c.Exp += amount
	fmt.Printf("✨ %s gagne %d points d'expérience ! (%d/%d)\n",
		c.Name, amount, c.Exp, c.ExpNext)

	if c.Exp >= c.ExpNext {
		c.LevelUp()
	}
}

// === Passage de niveau ===
func (c *Character) LevelUp() {
	c.Level++
	c.Exp -= c.ExpNext
	c.ExpNext += 10

	switch c.Class {
	case "Paladin":
		c.MaxHP += 15
		c.MaxMana += 5
		c.Atk += 3
		c.Def += 4
	case "Géant":
		c.MaxHP += 20
		c.Atk += 5
		c.Def += 5
	case "Mage":
		c.MaxHP += 8
		c.MaxMana += 15
		c.Atk += 2
		c.Def += 1
	case "Guérisseur":
		c.MaxHP += 10
		c.MaxMana += 12
		c.Atk += 1
		c.Def += 2
	}

	c.HP = c.MaxHP
	c.Mana = c.MaxMana

	fmt.Printf("\n🎉 %s passe au niveau %d !\n", c.Name, c.Level)
	fmt.Printf("Stats : HP %d | Mana %d | ATK %d | DEF %d\n",
		c.MaxHP, c.MaxMana, c.Atk, c.Def)
}

// === Gestion de l’inventaire ===
func (c *Character) AddItem(item string, qty int) {
	c.Inventory[item] += qty
	fmt.Printf("🧳 Vous obtenez %d x %s\n", qty, item)
}

func (c *Character) RemoveItem(item string, qty int) bool {
	if c.Inventory[item] < qty {
		return false
	}
	c.Inventory[item] -= qty
	if c.Inventory[item] == 0 {
		delete(c.Inventory, item)
	}
	return true
}

// === Utiliser un objet sur un joueur ===
func (c *Character) UseItem(item string, target *Character) {
	switch item {
	case "Potion de soin":
		if c.RemoveItem(item, 1) {
			heal := 50
			target.HP += heal
			if target.HP > target.MaxHP {
				target.HP = target.MaxHP
			}
			fmt.Printf("🍷 %s utilise une potion de soin → %s récupère %d PV (HP: %d/%d)\n",
				c.Name, target.Name, heal, target.HP, target.MaxHP)
		} else {
			fmt.Println("❌ Aucune potion de soin disponible.")
		}
	}
}

// === Utiliser un objet sur un monstre ===
func (c *Character) UseItemOnMonster(item string, target *monster.Monster) {
	switch item {
	case "Potion de poison":
		if c.RemoveItem(item, 1) {
			target.ApplyStatus("Poison", 3, 5)
		} else {
			fmt.Println("❌ Aucune potion de poison disponible.")
		}
	default:
		fmt.Println("❌ Objet inconnu :", item)
	}
}

// === Apprendre un nouveau sort ===
func (c *Character) LearnSkill(newSkill skills.Skill) {
	for _, s := range c.Skills {
		if s.Name == newSkill.Name {
			fmt.Println("❌ Vous connaissez déjà ce sort :", newSkill.Name)
			return
		}
	}
	c.Skills = append(c.Skills, newSkill)
	fmt.Println("✨ Nouveau sort appris :", newSkill.Name)
}

// === Utiliser un sort sur un monstre ===
func (c *Character) UseSkillOnMonster(skillName string, target *monster.Monster) {
	var s *skills.Skill
	for i := range c.Skills {
		if c.Skills[i].Name == skillName {
			s = &c.Skills[i]
			break
		}
	}
	if s == nil {
		fmt.Println("❌ Sort inconnu :", skillName)
		return
	}

	if c.Mana < s.ManaCost {
		fmt.Println("❌ Pas assez de mana pour lancer", s.Name)
		return
	}
	c.Mana -= s.ManaCost

	if s.IsHeal {
		healAmount := s.Power + (c.Mana / 10)
		c.HP += healAmount
		if c.HP > c.MaxHP {
			c.HP = c.MaxHP
		}
		fmt.Printf("✨ %s utilise %s → récupère %d PV (HP: %d/%d)\n",
			c.Name, s.Name, healAmount, c.HP, c.MaxHP)
		return
	}

	rawDamage := s.Power + (c.Atk / 2)
	if s.IsMagic {
		rawDamage = s.Power + (c.Mana / 5)
	}
	finalDamage := rawDamage - target.Def
	if finalDamage < 0 {
		finalDamage = 0
	}

	target.HP -= finalDamage
	if target.HP < 0 {
		target.HP = 0
	}

	fmt.Printf("🔥 %s utilise %s sur %s → %d dégâts (HP restants : %d/%d)\n",
		c.Name, s.Name, target.Name, finalDamage, target.HP, target.HPMax)
}

// Statuts
func (c *Character) ApplyStatus(name string, duration int, damage int) {
	for _, s := range c.Statuses {
		if s.Name == name {
			fmt.Printf("%s est déjà affecté par %s.\n", c.Name, name)
			return
		}
	}
	c.Statuses = append(c.Statuses, Status{Name: name, Duration: duration, Damage: damage})
	fmt.Printf("%s est maintenant affecté par %s (%d tours).\n", c.Name, name, duration)
}

func (c *Character) UpdateStatuses() {
	var remaining []Status
	for _, s := range c.Statuses {
		if s.Name == "Poison" {
			c.HP -= s.Damage
			fmt.Printf("☠️ %s subit %d dégâts de poison (HP: %d/%d)\n", c.Name, s.Damage, c.HP, c.MaxHP)
			if c.HP <= 0 {
				c.HP = 0
				fmt.Printf("💀 %s est mort à cause du poison !\n", c.Name)
				return
			}
		}
		s.Duration--
		if s.Duration > 0 {
			remaining = append(remaining, s)
		} else {
			fmt.Printf("%s n’est plus affecté par %s.\n", c.Name, s.Name)
		}
	}
	c.Statuses = remaining
}

// Recalculer les stats à partir de l’équipement
func (c *Character) RecalculateStatsFromEquipment() {
	hp, mana, atk, def, spd, crit := c.Equip.TotalStats()
	c.MaxHP += hp
	c.MaxMana += mana
	c.Atk += atk
	c.Def += def
	c.Spd += spd
	c.Crit += crit

	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
}

// Affichages
func (c Character) DisplaySummary() {
	fmt.Println("\n--- Résumé ---")
	fmt.Printf("Nom   : %s (%s)\n", c.Name, c.Gender)
	fmt.Println("Classe:", c.Class)
	fmt.Println("Niveau:", c.Level, "| XP:", c.Exp, "/", c.ExpNext)
}

func (c Character) DisplayFull() {
	fmt.Println("\n=== Informations du personnage ===")
	fmt.Printf("Nom    : %s (%s)\n", c.Name, c.Gender)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d | XP : %d/%d\n", c.Level, c.Exp, c.ExpNext)
	fmt.Printf("HP     : %d/%d\n", c.HP, c.MaxHP)
	fmt.Printf("Mana   : %d/%d\n", c.Mana, c.MaxMana)
	fmt.Printf("ATK    : %d | DEF: %d | SPD: %d | CRIT: %d%%\n", c.Atk, c.Def, c.Spd, c.Crit)
	fmt.Printf("Arme   : %s\n", c.Weapon)
	fmt.Printf("Or     : %d %s\n", c.Wallet.Amount, c.Wallet.Currency)

	if len(c.Skills) > 0 {
		fmt.Println("\n--- Compétences ---")
		for _, s := range c.Skills {
			fmt.Printf("- %s (Mana: %d) → %s\n", s.Name, s.ManaCost, s.Description)
		}
	} else {
		fmt.Println("Aucune compétence connue.")
	}

	if len(c.Statuses) > 0 {
		fmt.Println("\n--- Statuts ---")
		for _, s := range c.Statuses {
			fmt.Printf("- %s (%d tours restants)\n", s.Name, s.Duration)
		}
	}

	if len(c.Inventory) > 0 {
		fmt.Println("\n--- Inventaire ---")
		for item, qty := range c.Inventory {
			fmt.Printf("- %s x%d\n", item, qty)
		}
	}

	// 🔹 Affichage de l’équipement
	c.Equip.Display()
}

// === Charger un personnage depuis une sauvegarde ===
func LoadFromSave(state save.GameState) Character {
	c := Character{
		Name:      state.Name,
		Class:     state.Class,
		Gender:    "Inconnu", // pas stocké dans la sauvegarde actuelle
		Level:     1,         // valeur par défaut
		Exp:       0,
		ExpNext:   10,
		Wallet:    money.Money{Amount: state.Money, Currency: "Gold"},
		Inventory: map[string]int{},
		Equip:     equipment.EquipmentSet{},
	}

	// Inventaire sauvegardé (si présent)
	if state.Inventory != nil {
		c.Inventory = state.Inventory
	}

	// Stats de base selon la classe
	stats, ok := class.Classes[state.Class]
	if ok {
		c.MaxHP = stats.HP
		c.HP = c.MaxHP
		c.MaxMana = stats.Mana
		c.Mana = c.MaxMana
		c.Atk = stats.Atk
		c.Def = stats.Def
		c.Spd = stats.Spd
		c.Crit = stats.Crit
		c.Weapon = stats.Weapon
		c.Skills = skills.ClassSkills[state.Class]
	}

	// Recalcule les bonus des équipements (si tu sauvegardes l’équipement plus tard)
	c.RecalculateStatsFromEquipment()

	return c
}
